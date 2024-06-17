package repository

import (
	"context"

	"github.com/alex-v08/inventory-fw/internal/entity"
)

const (
	qryInsertUser = `INSERT INTO users (email, name, password) VALUES (?, ?, ?)`
	qrySelectUser = `SELECT id, email, name, password FROM users WHERE email = ?`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {

	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err

}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}

	err := r.db.GetContext(ctx, u, qrySelectUser, email)
	return u, err
}
