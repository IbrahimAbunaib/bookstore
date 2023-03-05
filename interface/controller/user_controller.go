package controller

import (
	"net/http"

	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/Domain/model"

	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(c Context) error {
	var params model.User

	if err := c.bind(&params); err != nil {
		return err
	}

	u, err := uc.userInteractor.Create(&params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
