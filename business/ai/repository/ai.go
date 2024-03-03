package repository

import (
	
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

func(r *AIRepository) OpenAI(content string) {

	prompt := struct {
		Input string `json:"input"`
	}{}

	prompt.Input = content
	// create the LLM
	llm, err := openai.NewChat(openai.WithModel(r.aiConfig.OpenAIModel))
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