package main

import (
	"log"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
	"fmt"
	"bytes"

	"document-scraping-with-ai/config"
	aiRepository "document-scraping-with-ai/business/ai/repository"
	// aiService "document-scraping-with-ai/business/ai/service"
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

	log.Println("\033[93mStarted. Press CTRL+C to quit.\033[0m")

	path := "./assets"

	entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    testRunPython()

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

func testRunPython() {

	path, pathErr := os.Getwd()
	if pathErr != nil {
	    fmt.Print(pathErr)
	}
	// handle err
	pythonPath := fmt.Sprintf("%s\\%s",path,"python\\pdf.py")
    fmt.Println("Run python")
    cmd := exec.Command("python", pythonPath)

    var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
	    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	    return
	}
	fmt.Println("Result: " + out.String())
}