package bootstrap

import (
	"context"
	"go-todoApp/api/controllers"
	"go-todoApp/api/middlewares"
	"go-todoApp/api/routes"
	"go-todoApp/libs"
	"go-todoApp/repository"
	"go-todoApp/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	middlewares.Module,
	routes.Module,
	libs.Module,
	services.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler libs.RequestHandler,
	routes routes.Routes,
	env libs.Env,
	logger libs.Logger,
	middleware middlewares.Middlewares,
	database libs.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")

			conn.SetMaxOpenConns(5)
			go func() {
				middleware.SetUp()
				routes.SetUp()
				handler.Gin.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
