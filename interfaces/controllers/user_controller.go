package controllers

import (
	"github.com/ryuseikurata/DDD/usecase"
	"github.com/ryuseikurata/DDD/interfaces/database"
	
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}