package ai

type Repository interface {
	OpenAI(content string) (string, error)
}