package application

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var db *sqlx.DB

var err error

func NewPostgreSQLDBConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	db, err = sqlx.Connect(os.Getenv("DB_DRIVER"), psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	createTableSchema, err := os.ReadFile("./database/tables/create.sql")
	if err = execSchema(ctx, db, string(createTableSchema)); err != nil {
		return nil
	}

	return db
}

func execSchema(ctx context.Context, db *sqlx.DB, schema string) (err error) {
	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("error execute schema %w", err)
	}

	return nil
}
