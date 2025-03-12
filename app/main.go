package main

import (
	"itv_go/config"
	"itv_go/database"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,     // Конфиг
			database.NewDatabase, // Теперь принимает *config.Config
		),
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{}
		}),
	)

	app.Run()
}
