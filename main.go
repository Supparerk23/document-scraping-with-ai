package main

import (
	"log"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
	"fmt"
	"strings"

	"document-scraping-with-ai/config"
	aiRepository "document-scraping-with-ai/business/ai/repository"
	// aiService "document-scraping-with-ai/business/ai/service"
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

	aiConfig := config.AIConfig()
	aiRepository := aiRepository.NewAIRepository(aiConfig)
	// aiService := aiService.NewAIService(aiRepository)

	pdfConfig := config.PdfConfig()

	cmd := exec.Command("sh", "-c", "./pipenv_script.sh")

	operationService := opService.NewPDFService(pdfConfig, cmd)

	log.Println("\033[93mStarted. Press CTRL+C to quit.\033[0m")

	err := operationService.ProcessPDF()

   	if err != nil {
        log.Fatal(err)
    }

	sourcePath := "./assets/raw"
	resultPath := "./assets/result"

	entries, err := os.ReadDir(sourcePath)
    if err != nil {
        log.Fatal(err)
    }


    for _, e := range entries {

    	fileDirectory := fmt.Sprintf("%s/%s",sourcePath,e.Name())

    	fmt.Println("ReadFile ->",fileDirectory)

    	b, err := os.ReadFile(fileDirectory) // just pass the file name
	    if err != nil {
	        fmt.Print(err)
	    }

	    str := string(b)

	    res := aiRepository.OpenAI(str)

	    fmt.Println("Res",res)

	    renameFile := strings.Replace(e.Name(), ".txt", ".json", -1)
	    resultFileDirectory := fmt.Sprintf("%s/%s",resultPath,renameFile)

	    if err = operationService.WriteResult(resultFileDirectory ,res); err != nil {
	    	fmt.Println("error write file",err)
	    }

    }
}