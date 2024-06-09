package service

func (s *service)ProcessAI(content string) (string, error) {
	res, err := s.aiRepo.OpenAI(content)
	if err != nil {
		return "", err
	}
	return res, nil
}