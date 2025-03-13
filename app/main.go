package main

import (
	"context"
	"fmt"
	"itv_go/config"
	"itv_go/database"
	_ "itv_go/docs"
	external "itv_go/external/ginapi"
	"itv_go/external/middleware"
	"itv_go/internal/repository/postgress"
	"itv_go/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func NewSwaggerHandler(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RunServer(lc fx.Lifecycle, router *gin.Engine, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := fmt.Sprintf(":%s", cfg.Port)

			go func() {
				if err := router.Run(port); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()

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
			middleware.NewAuthMiddleware,
			usecase.NewJwtUsecase,
			usecase.NewUserUsecase,
		),
		fx.Invoke(
			NewSwaggerHandler,
			RunServer,
			external.NewUserExternal,
			external.NewMovieExternal,
		),
	)

	app.Run()
}
