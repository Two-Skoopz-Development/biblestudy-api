package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	auth "github.com/Two-Skoopz-Development/biblestudy-api/internal/middleware"
	repository "github.com/Two-Skoopz-Development/biblestudy-api/internal/repositories"
	"github.com/Two-Skoopz-Development/biblestudy-api/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func getUser(conn *pgx.Conn, ctx context.Context, clerkId string) repository.GetUserByClerkIDRow {

	queries := repository.New(conn)
	user, err := queries.GetUserByClerkID(ctx, pgtype.Text{String: clerkId})
	if err != nil {
		log.Println(err)
	}
	log.Println(user)
	return user
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	clerkUser, ok := auth.GetUser(r.Context())
	if !ok || clerkUser == nil {
		utils.WriteError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	conn, err := utils.ConnectDB(ctx)
	if err != nil {
		log.Println(err)
		utils.WriteError(w, "Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	user := getUser(conn, ctx, clerkUser.ID)

	resp := fmt.Sprintf("User: %s", user.Fname.String)
	io.WriteString(w, resp)
}
