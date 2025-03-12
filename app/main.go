package main

import (
	"itv_go/config"
	"itv_go/database"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			database.NewDatabase,
		),
	)

	app.Run()
}
