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

func Testファクトリメソッド_アクティブユーザー生成(t *testing.T) {
	u := NewActiveUser()

	if !u.Active {
		t.Fatalf("active user expected")
	}
}

func TestDB接続_アクティブユーザー取得_トランザクションで独立(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()

	requireDBConnection(t, db)
	ensureSchema(t, db)

	withTx(t, db, func(tx *sql.Tx) {
		activeUser := NewUserBuilder().WithID("u-1").Build()
		inactiveUser := NewUserBuilder().WithID("u-2").WithInactive().Build()

		InsertUser(t, tx, activeUser)
		InsertUser(t, tx, inactiveUser)

		got, err := FindActiveUsers(tx)
		if err != nil {
			t.Fatalf("find active users: %v", err)
		}

		if len(got) != 1 || got[0].ID != "u-1" {
			t.Fatalf("unexpected result: %+v", got)
		}
	})
}
