package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"shop/internal/application"
)

func main() {
	application.Start()
}
