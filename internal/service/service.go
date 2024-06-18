package service

import (
	"context"

	"github.com/alex-v08/inventory-fw/internal/models"
	"github.com/alex-v08/inventory-fw/internal/repository"
)

// Service is an interface that wraps the basic CRUD operations for the User entity

//go:generate mockery --name Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
