package main

import (
	_config "sirclo/graphql/config"
	_graph "sirclo/graphql/delivery/controllers/graph"
	_userController "sirclo/graphql/delivery/controllers/user"

	_router "sirclo/graphql/delivery/router"
	_authRepo "sirclo/graphql/repository/auth"
	_bookRepo "sirclo/graphql/repository/book"
	_userRepo "sirclo/graphql/repository/user"
	_util "sirclo/graphql/util"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := _config.GetConfig()

	//initialize database connection based on given config
	db := _util.MysqlDriver(config)

	//initiate user model
	authRepo := _authRepo.New(db)
	userRepo := _userRepo.New(db)
	bookRepo := _bookRepo.New(db)

	//initiate user controller
	// authController := _authController.New(authRepo)
	userController := _userController.New(userRepo)

	//create echo http
	e := echo.New()
	client := _graph.NewResolver(authRepo, userRepo, bookRepo)
	srv := _router.NewGraphQLServer(client)
	//register API path and controller
	_router.RegisterPath(e, userController, srv)

	// run server
	address := fmt.Sprintf(":%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
