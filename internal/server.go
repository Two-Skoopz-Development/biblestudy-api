package api

import (
	"log"
	"net/http"
	"os"

	auth "github.com/Two-Skoopz-Development/biblestudy-api/internal/middleware"
	"github.com/Two-Skoopz-Development/biblestudy-api/internal/routes"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer() {

	r := routes.ConfigureRouter()

	r.MuxRouter.Use(mux.CORSMethodMiddleware(r.MuxRouter))
	r.MuxRouter.Use(auth.AuthMiddlware)
	apiKey := os.Getenv("CLERK_SECRET")
	// apiKey := "sk_test_bQ897iJxnDcVExE6qFPTM75ABlbYlBagS7ozxtQzp0"
	clerk.SetKey(apiKey)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins (for development, restrict in production)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)
	handler := corsHandler(r.MuxRouter)

	log.Fatal(http.ListenAndServe(":8080", handler), nil)

}
