package repository

import (
	"log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"

	"document-scraping-with-ai/model"
)

func(r *AIRepository) OpenAI(str string) (string, error) {

	llm, err := openai.New(openai.WithModel(r.aiConfig.OpenAIModel))
	if err != nil {
		return "", err
	}

	outputTemplate := model.ReturnTemplate{}

	template, err := json.Marshal(outputTemplate)
    if err != nil {
        return "", err
    }

	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem , "You are reading Fund Factsheet and helpful assistant designed to output JSON structured by assistant."),
		llms.TextParts(schema.ChatMessageTypeAI,string(template)),
		llms.TextParts(schema.ChatMessageTypeGeneric , str),
	}

	resp, err := llm.GenerateContent(ctx, content, llms.WithJSONMode())
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(resp.Choices[0].Content), &outputTemplate)
	if err != nil {
		panic(err)
	}

	fmt.Println(outputTemplate.FundCode)

	return resp.Choices[0].Content, nil

}

func showResponse(resp *llms.ContentResponse) string {
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}