package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"month02blogapi/config"
)

func main() {
	appConfig, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	db, err := sql.Open("mysql", appConfig.MySQLDSN)
	if err != nil {
		log.Fatalf("open mysql failed: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(appConfig.MaxOpenConns)
	db.SetMaxIdleConns(appConfig.MaxIdleConns)
	db.SetConnMaxLifetime(appConfig.ConnMaxLifetime())

	if err := db.Ping(); err != nil {
		log.Fatalf("ping mysql failed: %v", err)
	}

	fmt.Println("pool config:", appConfig.PoolSummary())

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("begin tx failed: %v", err)
	}

	title := "tx demo temporary post"
	_, err = tx.ExecContext(ctx, `
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`, title, "this row should be rolled back", "txdemo")
	if err != nil {
		_ = tx.Rollback()
		log.Fatalf("insert in tx failed: %v", err)
	}

	forcedErr := errors.New("simulate failure to demonstrate rollback")
	if forcedErr != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("rollback failed: %v", err)
		}
		fmt.Println("transaction rolled back:", forcedErr)
	} else {
		if err := tx.Commit(); err != nil {
			log.Fatalf("commit failed: %v", err)
		}
	}

	var count int
	if err := db.QueryRowContext(ctx, `SELECT COUNT(*) FROM posts WHERE title = ?`, title).Scan(&count); err != nil {
		log.Fatalf("verify rollback failed: %v", err)
	}

	fmt.Println("rows with temporary title after rollback:", count)
}