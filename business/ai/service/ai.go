package service

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"document-scraping-with-ai/model"
	"time"
)

func (s *service)ProcessAI(content string) (model.AIResponse, error) {

	var result model.AIResponse

	key := b64.StdEncoding.EncodeToString([]byte(content))
   
	found := true
		    
	data, err := s.redisClient.Get(key).Bytes()

	if err != nil {
		found = false
	}

	// result.ResultWithStruct
	if err := json.Unmarshal(data, &result); err != nil {
		found = false
	}

	if !found {

		res, err := s.aiRepo.OpenAI(content)
		if err != nil {
			return model.AIResponse{}, err
		}

		fmt.Println("-> api process")

		// res.ResultWithStruct
		dataByte, err := json.Marshal(res)
		
		err = s.redisClient.Set(key, dataByte, time.Hour*24).Err()
		if err != nil {
			fmt.Println("[WARNING] set befund to redis cache failed : %+v", err)
		}

		return res, nil

	}
	fmt.Println("-> user from cache")
	return result, nil
}