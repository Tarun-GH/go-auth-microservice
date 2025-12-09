package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Tarun-GH/go-rest-microservice/internal/models"
	"github.com/jackc/pgx/v5"
)

/* ---used in the beginning for hard coded column names---
func InsertUser(db *pgx.Conn, name, email, hash string) error {
	constQuery := "insert into users (name, email, password_hash) values ($1, $2, $3)"
	_, err := db.Exec(context.Background(), constQuery, name, email, hash)
	return err
}

	---this means PostgreSQL has a sequential and a default value
			is_identity = 'YES'
			column_default IS NOT NULL
*/

func InsertUser(db *pgx.Conn, name, email, hash string) error {
	query := `SELECT column_name 
	FROM information_schema.columns 
	WHERE table_name = $1 
	AND is_identity = 'NO' AND column_default IS NULL 
	ORDER BY ordinal_position;`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error querying column names: %w", err)
	}
	defer rows.Close()

	// colName := []string{name, email, hash} //given values of columns

	var columnName []string
	for rows.Next() {
		var colName string
		if err := rows.Scan(&colName); err != nil {
			return fmt.Errorf("error scanning column name: %w", err)
		}
		columnName = append(columnName, colName)
	}

	placeholders := make([]string, len(columnName))

	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	insertQuery := fmt.Sprintf("INSERT INTO users (%s) VALUES (%s)",
		strings.Join(columnName, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err = db.Exec(context.Background(), insertQuery, name, email, hash)
	return err
}

// To select a User by email
func GetUserByEmail(db *pgx.Conn, email string) (*models.User, error) {
	// ,err nhi use bcz only return 1 arg
	row := db.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1;", email)

	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
