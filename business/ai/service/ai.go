package service

import (
	"document-scraping-with-ai/model"
)

func (s *service)ProcessAI(content string) (model.AIResponse, error) {
	res, err := s.aiRepo.OpenAI(content)
	if err != nil {
		return model.AIResponse{}, err
	}
	return res, nil
}