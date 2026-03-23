package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./cmd/migrate up|down")
	}

	direction := os.Args[1]
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN is required")
	}

	migrationFile := filepath.Join("migrations", migrationName(direction))
	sqlBytes, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Fatalf("read migration file failed: %v", err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open mysql failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ping mysql failed: %v", err)
	}

	if _, err := db.Exec(string(sqlBytes)); err != nil {
		log.Fatalf("execute migration failed: %v", err)
	}

	fmt.Printf("migration %s applied: %s\n", direction, migrationFile)
}

func migrationName(direction string) string {
	switch direction {
	case "up":
		return "001_create_posts_table.up.sql"
	case "down":
		return "001_create_posts_table.down.sql"
	default:
		log.Fatalf("unsupported migration direction: %s", direction)
		return ""
	}
}
