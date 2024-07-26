package service

import (
	"github.com/go-redis/redis"
	"document-scraping-with-ai/business/ai"
)

type service struct {
	aiRepo ai.Repository
	redisClient *redis.Client
}


func NewAIService(aiRepo ai.Repository, redisClient *redis.Client) ai.Service {
	return &service{
		aiRepo: aiRepo,
		redisClient: redisClient,
	}
}