package services

import (
	"fmt"
	"io"
	"net/http"

	auth "github.com/Two-Skoopz-Development/biblestudy-api/internal/api/middleware"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := auth.GetUser(r.Context())
	if !ok || user == nil {
		io.WriteString(w, "Failed to get user")
		return
	}
	resp := fmt.Sprintf("User: %s", user.ID)
	io.WriteString(w, resp)
}
