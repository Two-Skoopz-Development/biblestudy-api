package utils

import (
	"log"
	"net/http"

	"github.com/Two-Skoopz-Development/biblestudy-api/internal/responses"
)

func WriteError(w http.ResponseWriter, errorMsg string, code int) {
	payload, err := responses.GenerateStatusMessage(errorMsg)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}
