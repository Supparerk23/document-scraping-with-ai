package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"fmt"
)

// initialise to load environment variable from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println("\033[93mStarted. Press CTRL+C to quit.\033[0m")
	run()
}

// call the LLM and return the response
func run() {

	path := "./assets"

	entries, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }
 
    for _, e := range entries {

    	fileDirectory := fmt.Sprintf("%s/%s",path,e.Name())

    	fmt.Println("ReadFile ->",fileDirectory)

    	b, err := os.ReadFile(fileDirectory) // just pass the file name
	    if err != nil {
	        fmt.Print(err)
	    }

	    str := string(b)

	    aiCaller(str)
        
    }
}

func aiCaller(content string) {

	prompt := struct {
		Input string `json:"input"`
	}{}

	prompt.Input = content
	// create the LLM
	llm, err := openai.NewChat(openai.WithModel(os.Getenv("OPENAI_MODEL")))
	if err != nil {
		fmt.Print(err)
	}

	chatmsg := []schema.ChatMessage{
		// schema.SystemChatMessage{Content: "Hello, I am a friendly AI assistant."},
		schema.HumanChatMessage{Content: prompt.Input},
	}

	fmt.Println("Processing by AI")

	aimsg, err := llm.Call(context.Background(), chatmsg)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Process Done")
	fmt.Println(aimsg.GetContent())

}