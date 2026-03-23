package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN is required, example: root:password@tcp(127.0.0.1:3306)/mysql?parseTime=true")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open mysql failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping mysql failed: %v", err)
	}

	fmt.Println("mysql connection ok")

	var version string
	if err := db.QueryRow("SELECT VERSION()").Scan(&version); err != nil {
		log.Fatalf("query mysql version failed: %v", err)
	}
	fmt.Println("mysql version:", version)

	var one int
	if err := db.QueryRow("SELECT 1").Scan(&one); err != nil {
		log.Fatalf("simple query failed: %v", err)
	}
	fmt.Println("simple query result:", one)
}
