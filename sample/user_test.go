package sample

import (
	"database/sql"
	"testing"
)

func Testユーザービルダー_非アクティブ生成(t *testing.T) {
	u := NewUserBuilder().WithID("u-999").WithInactive().Build()

	if u.Active {
		t.Fatalf("inactive user expected")
	}
}

func TestDB接続_アクティブユーザー取得_トランザクションで独立(t *testing.T) {
	db, err := openTestDB()
	if err != nil {
		t.Skipf("TEST_DSN が未設定のためスキップ: %v", err)
	}
	defer db.Close()

	if err := withTx(db, func(tx *sql.Tx) error {
		activeUser := NewUserBuilder().WithID("u-1").Build()
		inactiveUser := NewUserBuilder().WithID("u-2").WithInactive().Build()

		if err := InsertUser(tx, activeUser); err != nil {
			return err
		}
		if err := InsertUser(tx, inactiveUser); err != nil {
			return err
		}

		// ここにSUT呼び出しを置く想定（例: FindActiveUsers）
		// got, err := FindActiveUsers(tx)
		// ...

		return nil
	}); err != nil {
		t.Fatalf("tx test failed: %v", err)
	}
}
