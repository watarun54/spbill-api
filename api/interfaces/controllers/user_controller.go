package controllers

import (
	"strconv"

	"github.com/watarun54/serverless-skill-manager/server/domain"
	"github.com/watarun54/serverless-skill-manager/server/interfaces/database"
	"github.com/watarun54/serverless-skill-manager/server/usecase"
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

func (controller *UserController) GetMe(c Context) (err error) {
	uid := userIDFromToken(c)
	user, err := controller.Interactor.UserById(uid)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(user))
	return
}

func (controller *UserController) UpdateMe(c Context) (err error) {
	uid := userIDFromToken(c)
	u := domain.User{
		ID: uid,
	}
	c.Bind(&u)
	user, err := controller.Interactor.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, NewResponse(user))
	return
}

func (controller *UserController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(user))
	return
}

func (controller *UserController) Index(c Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(users))
	return
}

func (controller *UserController) Create(c Context) (err error) {
	uForm := domain.UserForm{}
	c.Bind(&uForm)
	uForm.HashedPassword = generateHash(uForm.Email, uForm.Password)
	u := controller.Interactor.ConvertUserFormToUser(uForm)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, NewResponse(user))
	return
}

func (controller *UserController) Save(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := domain.User{
		ID: id,
	}
	c.Bind(&u)
	user, err := controller.Interactor.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, NewResponse(user))
	return
}

func (controller *UserController) Delete(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{
		ID: id,
	}
	err = controller.Interactor.DeleteById(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(user))
	return
}
