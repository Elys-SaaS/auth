package interfaces

import "github.com/Elys-SaaS/auth/model"

type UserService interface {
	GetByID(id int) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
