package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/Two-Skoopz-Development/biblestudy-api/internal/utils"
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

			utils.WriteError(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

	return clerkhttp.RequireHeaderAuthorization()(handler)

}
