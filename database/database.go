package database

import (
	"context"
	"fmt"
	"github.com/alex-v08/inventory-fw/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)

	db, err := sqlx.ConnectContext(ctx, "mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Optional: Check if the connection is successful
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
