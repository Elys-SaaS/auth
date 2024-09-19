package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	auth := v1.Group("/auth")
	auth.POST("/signup", h.SignUp)
	auth.POST("/signin", h.SignIn)
	auth.GET("/refresh", h.Refresh)
}
