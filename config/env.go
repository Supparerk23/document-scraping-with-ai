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

func PdfConfig() model.PdfConfig {
	return model.PdfConfig{
		PathLocation: os.Getenv("PDF_PATH_LOCATION"),
	}
}