package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Two-Skoopz-Development/biblestudy-api/internal/api/responses"
	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

func GetClaims(ctx context.Context) (*clerk.SessionClaims, bool) {
	return clerk.SessionClaimsFromContext(ctx)
}

func GetUser(ctx context.Context) (*clerk.User, bool) {
	var claims *clerk.SessionClaims
	var ok bool
	var err error
	var usr *clerk.User

	if claims, ok = GetClaims(ctx); !ok {
		return &clerk.User{}, false
	}

	if usr, err = user.Get(ctx, claims.Subject); err != nil {
		return &clerk.User{}, false
	}

	return usr, true

}

func AuthMiddlware(next http.Handler) http.Handler {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, ok := clerk.SessionClaimsFromContext(r.Context())
		err := claims.ValidateWithLeeway(time.Now(), 0)

		if !ok || err != nil {
			response := responses.StatusMessage{Message: "unauthoraized"}
			payload, err := json.Marshal(response)
			if err != nil {
				log.Println("Error: error marshalling json")
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(payload)
			return
		}

		next.ServeHTTP(w, r)
	})

	return clerkhttp.RequireHeaderAuthorization()(handler)

}
