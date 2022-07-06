package controllers

import (
	"strconv"

	"github.com/watarun54/spbill-api/server/domain"
	"github.com/watarun54/spbill-api/server/interfaces/database"
	"github.com/watarun54/spbill-api/server/usecase"
)

type RoomController struct {
	Interactor           usecase.RoomInteractor
	BillInteractor       usecase.BillInteractor
	RoomMemberInteractor usecase.RoomMemberInteractor
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
		BillInteractor: usecase.BillInteractor{
			BillRepository: &database.BillRepository{
				SqlHandler: sqlHandler,
			},
			RoomRepository: &database.RoomRepository{
				SqlHandler: sqlHandler,
			},
			RoomMemberRepository: &database.RoomMemberRepository{
				SqlHandler: sqlHandler,
			},
		},
		RoomMemberInteractor: usecase.RoomMemberInteractor{
			RoomMemberRepository: &database.RoomMemberRepository{
				SqlHandler: sqlHandler,
			},
			BillRepository: &database.BillRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RoomController) FindByUUID(c Context) (err error) {
	room, err := controller.Interactor.FindByUUID(c.Param("uuid"))
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
	rForm := domain.RoomForm{}
	c.Bind(&rForm)
	r, err := controller.Interactor.ConvertRoomFormToRoom(rForm)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	room, err := controller.Interactor.Add(r)
	room.RoomMembers = make([]domain.RoomMember, 0)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}

func (controller *RoomController) Update(c Context) (err error) {
	rForm := domain.RoomForm{}
	rForm.UUID = c.Param("uuid")
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
	err = controller.Interactor.DeleteByUUID(domain.Room{UUID: c.Param("uuid")})
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, nil)
	return
}

func (controller *RoomController) FetchBills(c Context) (err error) {
	room, err := controller.Interactor.FindByUUID(c.Param("uuid"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	bill := domain.Bill{
		RoomID: room.ID,
	}
	c.Bind(&bill)
	bills, err := controller.BillInteractor.Bills(bill)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, bills)
	return
}

func (controller *RoomController) UserPayments(c Context) (err error) {
	room, err := controller.Interactor.FindByUUID(c.Param("uuid"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	bill := domain.Bill{
		RoomID: room.ID,
	}
	c.Bind(&bill)
	userPayments, err := controller.BillInteractor.UserPayments(bill)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, userPayments)
	return
}

func (controller *RoomController) AddBill(c Context) (err error) {
	room, err := controller.Interactor.FindByUUID(c.Param("uuid"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	bForm := domain.BillForm{RoomID: room.ID}
	c.Bind(&bForm)
	b, err := controller.BillInteractor.ConvertBillFormToBill(bForm)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	_, err = controller.BillInteractor.Add(b)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	newRoom, err := controller.Interactor.Room(room)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, newRoom)
	return
}

func (controller *RoomController) AddMember(c Context) (err error) {
	room, err := controller.Interactor.FindByUUID(c.Param("uuid"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	roomMember := domain.RoomMember{
		RoomID: room.ID,
	}
	c.Bind(&roomMember)
	_, err = controller.RoomMemberInteractor.Add(roomMember)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	newRoom, err := controller.Interactor.Room(room)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, newRoom)
	return
}

func (controller *RoomController) DeleteMember(c Context) (err error) {
	memberId, _ := strconv.Atoi(c.Param("member_id"))
	err = controller.RoomMemberInteractor.Delete(domain.RoomMember{ID: memberId})
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	room, err := controller.Interactor.FindByUUID(c.Param("uuid"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, room)
	return
}
