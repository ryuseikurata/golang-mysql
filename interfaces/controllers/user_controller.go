package controllers

import (
	"github.com/ryuseikurata/DDD/usecase"
	"github.com/ryuseikurata/DDD/domain"
	"github.com/ryuseikurata/DDD/interfaces/database"
	"strconv"
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

func (controller *UserController) Create(c Context) {
	user := domain.User{}
	c.Bind(&user)
	err := controller.Interactor.Add(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return 
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}