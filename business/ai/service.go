package ai

import (
	"document-scraping-with-ai/model"
)

type Service interface {
	ProcessAI(orginalName string, content string) (model.AIResponse, error)
}