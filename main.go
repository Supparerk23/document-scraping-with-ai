package main

import (
	"log"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
	"fmt"
	"strings"
	"time"

	"document-scraping-with-ai/config"
	aiRepository "document-scraping-with-ai/business/ai/repository"
	aiService "document-scraping-with-ai/business/ai/service"
	opService "document-scraping-with-ai/business/operation/service"
)

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	start := time.Now()

	// Init Redis connection
	config.InitRedis()
	redisClient := config.GetRedis()
	defer config.CloseRedis()

	aiConfig := config.AIConfig()
	aiRepository := aiRepository.NewAIRepository(aiConfig)
	aiService := aiService.NewAIService(aiRepository,redisClient)

	pdfConfig := config.PdfConfig()

	cmd := exec.Command("sh", "-c", "./pipenv_script.sh")

	operationService := opService.NewPDFService(pdfConfig, cmd)

	log.Println("\x1b[46mStarted. Press CTRL+C to quit.\033[0m")

	err := operationService.ProcessPDF()

   	if err != nil {
        log.Fatal(err)
    }

	sourcePath := "./assets"
	resultPath := "./assets/result"

	entries, err := os.ReadDir(sourcePath)
    if err != nil {
        log.Fatal(err)
    }

    for _, e := range entries {

    	if strings.Contains(e.Name(), "raw") {

	    	fileDirectory := fmt.Sprintf("%s/%s",sourcePath,e.Name())

	    	originalFileName := strings.Replace(e.Name(), "_raw.txt", "", -1)

	    	fmt.Println("Process File >",fileDirectory)

	    	b, err := os.ReadFile(fileDirectory)
		    if err != nil {
		        fmt.Print(err)
		    }

		    str := string(b)

		    if len(str) == 0 {
		    	fmt.Println("File : ",fileDirectory,"is empty.")
		    	return
		    }

		    res, err := aiService.ProcessAI(originalFileName,str)
		    if err != nil {
		    	fmt.Print(err)
		    }

		    // this return should save data to database
			// fmt.Println("Save to database",res.ResultWithStruct)

		    renameFileRaw := strings.Replace(e.Name(), ".txt", ".json", -1)

		    // fmt.Println("Save File >",renameFileRaw)

		    if err = operationService.WriteResult(fmt.Sprintf("%s/%s",resultPath,renameFileRaw) ,res.RawResult); err != nil {
		    	fmt.Println("error write file",err)
		    }

		    e := os.Remove(fileDirectory) 
		    if e != nil { 
		        log.Fatal(e) 
		    }

		    // fmt.Println("Del File >",fileDirectory)

	    }

    }

    elapsed := time.Since(start)
    // fmt.Sprintf("\x1b[46mDone.\033[0m - complete, Elapsed time %s",elapsed)
    log.Printf("\x1b[46mDone.\033[0m - complete, Elapsed time %s",elapsed)
}