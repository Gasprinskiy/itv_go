package main

import (
	"context"
	"fmt"
	"itv_go/config"
	"itv_go/database"
	_ "itv_go/docs"
	external "itv_go/external/ginapi"
	"itv_go/internal/repository/postgress"
	"itv_go/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func NewSwaggerHandler(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RunServer(lc fx.Lifecycle, router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := router.Run(":8080"); err != nil {
					panic(err) // Если сервер упадет — сразу краш
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Server is shutting down...")
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			gin.New,
			config.NewConfig,
			database.NewDatabase,
			postgress.NewUserRepository,
			postgress.NewMovieRepository,
			usecase.NewJwtUsecase,
			usecase.NewUserUsecase,
		),
		fx.Invoke(
			NewSwaggerHandler,
			RunServer,
			external.NewUserExternal,
		),
	)

	app.Run()
}
