package repository

import (
	"log"
	"context"
	"fmt"
	"encoding/json"
	
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

func(r *AIRepository) OpenAI(str string) string{

	fmt.Println("----------------------->")

	// prompt := struct {
	// 	Input string `json:"input"`
	// }{}

	// prompt.Input = str
	// create the LLM
	// llm, err := openai.NewChat(openai.WithModel(r.aiConfig.OpenAIModel))
	llm, err := openai.New(openai.WithModel(r.aiConfig.OpenAIModel))
	if err != nil {
		fmt.Print(err)
	}

	// chatmsg := []schema.ChatMessage{
	// 	// schema.SystemChatMessage{Content: "Hello, I am a friendly AI assistant."},
	// 	// schema.HumanChatMessage{Content: prompt.Input},
	// 	// schema.GenericChatMessage{
	// 	// 	{
	// 	// 		Content: "You are reading Fund Factsheet and helpful assistant designed to output JSON structured by assistant.",
	// 	// 		Role : "system",
	// 	// 	},
	// 	// },
	// 	// schema.GenericChatMessage{
	// 	// 	{
	// 	// 		Content: "{'fund_code':','mutual_fund_fee':{'management':{'max':','actual':'},'total':{'max':','actual':'}},'unit_holder_fee':{'sale':{'max':','actual':'},'buy_back':{'max':','actual':'},'switch_in':{'max':','actual':'},'switch_out':{'max':','actual':'}}}",
	// 	// 		Role : "assistant",
	// 	// 	},
	// 	// },
	// 	// schema.GenericChatMessage{
	// 	// 	{
	// 	// 		Content: prompt.Input,
	// 	// 		Role : "user",
	// 	// 	},
	// 	// },
	// 	schema.GenericChatMessage{
	// 		Content: "You are reading Fund Factsheet and helpful assistant designed to output JSON structured by assistant.",
	// 		Role : "system",
	// 	},
	// 	schema.GenericChatMessage{
	// 		Content: "{'fund_code':','mutual_fund_fee':{'management':{'max':','actual':'},'total':{'max':','actual':'}},'unit_holder_fee':{'sale':{'max':','actual':'},'buy_back':{'max':','actual':'},'switch_in':{'max':','actual':'},'switch_out':{'max':','actual':'}}}",
	// 		Role : "assistant",
	// 	},
	// 	schema.GenericChatMessage{
	// 		Content: prompt.Input,
	// 		Role : "user",
	// 	},
	// }


	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem , "You are reading Fund Factsheet and helpful assistant designed to output JSON structured by assistant."),
		llms.TextParts(schema.ChatMessageTypeAI, "{'fund_code':','mutual_fund_fee':{'management':{'max':','actual':'},'total':{'max':','actual':'}},'unit_holder_fee':{'sale':{'max':','actual':'},'buy_back':{'max':','actual':'},'switch_in':{'max':','actual':'},'switch_out':{'max':','actual':'}}}"),
		llms.TextParts(schema.ChatMessageTypeGeneric , str),
	}

	fmt.Println("Processing by AI")

	// aimsg, err := llm.Call(context.Background(), chatmsg, llms.WithJSONMode())
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// fmt.Println("Process Done")
	// fmt.Println(aimsg.GetContent())

	// if _, err := llm.GenerateContent(ctx, content,
	// 	// llms.WithMaxTokens(1024),
	// 	llms.WithJSONMode(),
	// 	llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
	// 		fmt.Print(string(chunk))
	// 		return nil
	// 	})); err != nil {
	// 	log.Fatal(err)
	// }

	resp, err := llm.GenerateContent(ctx, content, llms.WithJSONMode())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Process Done")
	// fmt.Println("Initial response:", showResponse(resp))
	// fmt.Println(resp.Choices[0].Content)

	return resp.Choices[0].Content

}

func showResponse(resp *llms.ContentResponse) string {
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}