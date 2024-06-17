package main

import (
	"context"

	"github.com/alex-v08/inventory-fw/database"
	"github.com/alex-v08/inventory-fw/internal/repository"
	"github.com/alex-v08/inventory-fw/internal/service"
	"github.com/alex-v08/inventory-fw/settings"

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
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
				if err != nil {
					panic(err)
				}
				u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")
				if err != nil {
					panic(err)
				}
				if u.Name != "myname" {
					panic("name does not match")
				}
			},
		),
	)

	app.Run()

}
