package controllers

import (
	"strconv"

	"github.com/watarun54/serverless-skill-manager/server/domain"
	"github.com/watarun54/serverless-skill-manager/server/interfaces/database"
	"github.com/watarun54/serverless-skill-manager/server/usecase"
)

type RoomController struct {
	Interactor usecase.RoomInteractor
}

func NewRoomController(sqlHandler database.SqlHandler) *RoomController {
	return &RoomController{
		Interactor: usecase.RoomInteractor{
			RoomRepository: &database.RoomRepository{
				SqlHandler: sqlHandler,
			},
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RoomController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := domain.Room{
		ID: id,
	}
	room, err := controller.Interactor.Room(r)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}

func (controller *RoomController) Index(c Context) (err error) {
	rooms, err := controller.Interactor.Rooms()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, rooms)
	return
}

func (controller *RoomController) Create(c Context) (err error) {
	r := domain.Room{}
	c.Bind(&r)
	room, err := controller.Interactor.Add(r)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}

func (controller *RoomController) Update(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	rForm := domain.RoomForm{}
	rForm.ID = id
	c.Bind(&rForm)
	r, err := controller.Interactor.ConvertRoomFormToRoom(rForm)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	room, err := controller.Interactor.Update(r)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}

func (controller *RoomController) Delete(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	room := domain.Room{
		ID: id,
	}
	c.Bind(&room)
	err = controller.Interactor.DeleteById(room)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}
