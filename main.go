package main

import (
	"log"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
	"fmt"

	"document-scraping-with-ai/config"
	aiRepository "document-scraping-with-ai/business/ai/repository"
	// aiService "document-scraping-with-ai/business/ai/service"
	pdfService "document-scraping-with-ai/business/pdf/service"
)

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	ai := false
	aiConfig := config.AIConfig()
	aiRepository := aiRepository.NewAIRepository(aiConfig)
	// aiService := aiService.NewAIService(aiRepository)

	pdfConfig := config.PdfConfig()

	cmd := exec.Command("sh", "-c", "./pipenv_script.sh")

	pdfService := pdfService.NewPDFService(pdfConfig, cmd)

	log.Println("\033[93mStarted. Press CTRL+C to quit.\033[0m")

	path := "./assets"

	entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

   	pdfService.ProcessPDF()

    if ai {

    for _, e := range entries {

    	fileDirectory := fmt.Sprintf("%s/%s",path,e.Name())

    	fmt.Println("ReadFile ->",fileDirectory)

    	b, err := os.ReadFile(fileDirectory) // just pass the file name
	    if err != nil {
	        fmt.Print(err)
	    }

	    str := string(b)

	    aiRepository.OpenAI(str)
        
    }

    }
}