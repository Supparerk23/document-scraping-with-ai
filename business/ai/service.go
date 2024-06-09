package ai

type Service interface {
	ProcessAI(content string) (string, error)
}