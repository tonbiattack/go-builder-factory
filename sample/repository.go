package sample

import (
	"database/sql"
	"fmt"
)

// FindActiveUsers はアクティブなユーザーを取得するSUTの例。
func FindActiveUsers(tx *sql.Tx) ([]User, error) {
	rows, err := tx.Query(
		`SELECT id, name, email, active FROM users WHERE active = ? ORDER BY id`,
		true,
	)
	if err != nil {
		return nil, fmt.Errorf("query active users: %w", err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Active); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}
