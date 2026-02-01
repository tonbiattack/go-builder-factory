package sample

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// User はドメインモデルの例。
type User struct {
	ID     string
	Name   string
	Email  string
	Active bool
}

// NewActiveUser は固定パターンのファクトリメソッド例。
func NewActiveUser() User {
	return User{
		ID:     "u-001",
		Name:   "Alice",
		Email:  "alice@example.com",
		Active: true,
	}
}

// UserBuilder は差分指定で生成できるビルダー。
type UserBuilder struct {
	user User
}

// NewUserBuilder は標準値で初期化したビルダーを返す。
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: User{
		ID:     "u-default",
		Name:   "Default",
		Email:  "default@example.com",
		Active: true,
	}}
}

func (b *UserBuilder) WithID(id string) *UserBuilder {
	b.user.ID = id
	return b
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithInactive() *UserBuilder {
	b.user.Active = false
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

// withTx はトランザクションを開始し、終了時にロールバックする。
// テスト用の使い捨てデータを作るためのヘルパー。
func withTx(db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	return fn(tx)
}

// InsertUser はテスト用データのInserter。
func InsertUser(tx *sql.Tx, u User) error {
	_, err := tx.Exec(`INSERT INTO users(id, name, email, active) VALUES (?, ?, ?, ?)`
		, u.ID, u.Name, u.Email, u.Active,
	)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}

// openTestDB はテスト用DBを開く。
// 例: TEST_DSN="root:password@tcp(localhost:3306)/test_db?parseTime=true"
func openTestDB() (*sql.DB, error) {
	dsn := os.Getenv("TEST_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("TEST_DSN is required")
	}

	return sql.Open("mysql", dsn)
}
