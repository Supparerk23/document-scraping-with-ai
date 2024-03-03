package repository

import (
	"document-scraping-with-ai/business/ai"
	"document-scraping-with-ai/model"
)

type AIRepository struct {
	aiConfig model.AIConfig
}

func NewAIRepository(aiConfig model.AIConfig) ai.Repository {
	return &AIRepository{
		aiConfig: aiConfig,
	}
}