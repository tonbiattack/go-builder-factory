package sample

import (
	"database/sql"
	"testing"
)

func TestUserBuilder_Inactive(t *testing.T) {
	t.Run("ユーザービルダー_非アクティブ生成", func(t *testing.T) {
		u := NewUserBuilder().WithID("u-999").WithInactive().Build()

		if u.Active {
			t.Fatalf("inactive user expected")
		}
	})
}

func TestFactoryMethod_ActiveUser(t *testing.T) {
	t.Run("ファクトリメソッド_アクティブユーザー生成", func(t *testing.T) {
		u := NewActiveUser()

		if !u.Active {
			t.Fatalf("active user expected")
		}
	})
}

func TestFindActiveUsers_WithTransaction(t *testing.T) {
	t.Run("DB接続_アクティブユーザー取得_トランザクションで独立", func(t *testing.T) {
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
	})
}
