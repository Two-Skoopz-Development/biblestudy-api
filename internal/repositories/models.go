// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID          int32       `json:"id"`
	ClerkID     pgtype.Text `json:"clerk_id"`
	Fname       pgtype.Text `json:"fname"`
	Lname       pgtype.Text `json:"lname"`
	Email       pgtype.Text `json:"email"`
	PhoneNumber pgtype.Text `json:"phone_number"`
	BibleID     pgtype.Int4 `json:"bible_id"`
	IsAdmin     pgtype.Bool `json:"is_admin"`
}
