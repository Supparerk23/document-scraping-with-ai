package ai

import (
	"document-scraping-with-ai/model"
)

type Repository interface {
	OpenAI(content string) (model.AIResponse, error)
}