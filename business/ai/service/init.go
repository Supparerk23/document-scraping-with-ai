package service

import (
	"document-scraping-with-ai/business/ai"
)

type service struct {
	aiRepo ai.Repository
}


func NewAIService(aiRepo ai.Repository) ai.Service {
	return &service{
		aiRepo: aiRepo,
	}
}