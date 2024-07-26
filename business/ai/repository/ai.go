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
	"document-scraping-with-ai/constant"
)

func(r *AIRepository) OpenAI(str string) (model.AIResponse, error) {

	output := model.AIResponse{}

	llm, err := openai.New(openai.WithModel(r.aiConfig.OpenAIModel))
	if err != nil {
		return output, err
	}

	outputTemplate := model.ReturnTemplate{}

	template, err := json.Marshal(outputTemplate)
    if err != nil {
        return output, err
    }

	ctx := context.Background()

	SystemPrompt :=  fmt.Sprintf(constant.SystemPrompt,string(template))

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem ,SystemPrompt),
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

	output.ResultWithStruct = outputTemplate
	output.RawResult = resp.Choices[0].Content

	return output, nil

}

func showResponse(resp *llms.ContentResponse) string {
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}