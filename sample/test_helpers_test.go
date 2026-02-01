package sample

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()

	dsn := os.Getenv("TEST_DSN")
	if dsn == "" {
		t.Skip("TEST_DSN が未設定のためスキップ")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}

	return db
}

func ensureSchema(t *testing.T, db *sql.DB) {
	t.Helper()

	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(64) PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	active BOOLEAN NOT NULL
);
`)
	if err != nil {
		t.Fatalf("create table: %v", err)
	}
}

func withTx(t *testing.T, db *sql.DB, fn func(tx *sql.Tx)) {
	t.Helper()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	fn(tx)
}

func InsertUser(t *testing.T, tx *sql.Tx, u User) {
	t.Helper()

	_, err := tx.Exec(
		`INSERT INTO users(id, name, email, active) VALUES (?, ?, ?, ?)`,
		u.ID, u.Name, u.Email, u.Active,
	)
	if err != nil {
		t.Fatalf("insert user: %v", err)
	}
}

func requireDBConnection(t *testing.T, db *sql.DB) {
	t.Helper()

	if err := db.Ping(); err != nil {
		t.Fatalf("ping db: %v", err)
	}
}

func formatTestDSN(dbName string) string {
	return fmt.Sprintf("root:password@tcp(localhost:3306)/%s?parseTime=true", dbName)
}
