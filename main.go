package main

import (
	"context"

	"github.com/alex-v08/inventory-fw/database"
	"github.com/alex-v08/inventory-fw/internal/repository"
	"github.com/alex-v08/inventory-fw/internal/service"
	"github.com/alex-v08/inventory-fw/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),

		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)

				}
			},
		),
	)

	app.Run()

}
