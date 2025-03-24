package utils

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", username, password, db_name, db_host, db_port)
	//NEW TEST
	return pgx.Connect(ctx, connectionString)

	//OLD SAVE
	// // Open the database connection
	// db, err := sql.Open("postgres", connectionString)
	// if err != nil {
	// 	return nil, err
	// }

	// // Check if the connection is successful
	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }

	// // fmt.Println("Connected to PostgreSQL!")
	// return db, nil

}
