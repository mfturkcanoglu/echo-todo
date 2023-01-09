package server

import (
	"github.com/mfturkcanoglu/echo-todo/pkg/client"
	"github.com/mfturkcanoglu/echo-todo/pkg/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var (
	logger, _                    = zap.NewDevelopment()
	psqlClient     client.Client = client.NewPostgresClient(logger)
	todoController               = controller.NewTodoController(psqlClient.Db(), logger)
)

func Server() *echo.Echo {
	defer logger.Sync()
	defer psqlClient.Close()

	e := echo.New()
	routes(e)
	e.Logger.Fatal(e.Start(":8080"))
	return e
}

// Add routes to echo
func routes(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	// todo endpoints
	e.GET("api/todo/:id", todoController.Get)
	e.GET("api/todo", todoController.GetAll)
	e.POST("api/todo", todoController.Post)
}
