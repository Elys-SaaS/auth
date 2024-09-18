package handler

import "github.com/Elys-SaaS/auth/interfaces"

type Handler struct {
	userService interfaces.UserService
}

func NewHandler(userService interfaces.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}
