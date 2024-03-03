package config

import (
	"os"
	"document-scraping-with-ai/model"
)

func AIConfig() model.AIConfig {
	return model.AIConfig{
		OpenAIApiKey: os.Getenv("OPENAI_API_KEY"),
		OpenAIModel: os.Getenv("OPENAI_MODEL"),
	}
}