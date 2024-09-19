package main

import (
	"github.com/Elys-SaaS/auth/db"
	"github.com/Elys-SaaS/auth/handler"
	"github.com/Elys-SaaS/auth/router"
	"github.com/Elys-SaaS/auth/services"
)

func main() {
	r := router.New()
	d := db.New()

	v1 := r.Group("/api")
	db.AutoMigrate(d)
	us := services.NewUserService(d)
	h := handler.NewHandler(us)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
