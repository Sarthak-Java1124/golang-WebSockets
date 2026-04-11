package users

import (
	"context"
	"database/sql"
	"log"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
type repository struct {
	db DBTX
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertedId int
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertedId)
	if err != nil {
		log.Printf("query error: %v", err)
		return nil, err
	}
	user.ID = int64(lastInsertedId)
	return user, nil
}

func (r *repository) GetUsersByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := "SELECT id,email,username,password FROM users  WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		log.Println("The error in getting users is : ", err)
		return nil, err
	}
	return &user, err

}
