package responses

import (
	"encoding/json"
)

type StatusMessage struct {
	Message string `json:"message"`
}

func GenerateStatusMessage(msg string) ([]byte, error) {
	statusMessage := StatusMessage{Message: msg}
	return json.Marshal(statusMessage)
}
