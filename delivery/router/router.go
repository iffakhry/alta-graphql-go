package router

import (
	_userController "sirclo/graphql/delivery/controllers/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, userController *_userController.UserController, srv *handler.Server) {
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	// Login
	// e.POST("/login", authController.Login())
	// e.POST("/persons", userController.Create())
	e.GET("/users", userController.Get())
	// e.POST("/books", bookController.Create())
	// e.GET("/books", bookController.Get())
	{
		e.Use(middleware.CORSWithConfig((middleware.CORSConfig{})))

		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		// For Subscriptions
		// e.GET("/subscriptions", func(c echo.Context) error {
		// 	srv.ServeHTTP(c.Response(), c.Request())
		// 	return nil
		// })

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

}
