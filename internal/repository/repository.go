package repository

import (
	"context"

	"github.com/alex-v08/inventory-fw/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is an interface that wraps the basic CRUD operations for the User entity
//
//go:generate mockery --name Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
