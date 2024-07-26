package ai

import (
	"document-scraping-with-ai/model"
)

type Service interface {
	ProcessAI(content string) (model.AIResponse, error)
}