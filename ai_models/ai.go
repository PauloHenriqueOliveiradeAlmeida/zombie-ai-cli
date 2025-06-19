package ai_models

type AI interface {
	GetResponse(prompt string, maxTokens int) (string, error)
}
