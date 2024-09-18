package main

import (
	"github.com/Elys-SaaS/auth/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @host 127.0.0.1:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := router.New()

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
