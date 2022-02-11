package controllers

import (
	"strconv"

	"github.com/watarun54/serverless-skill-manager/server/domain"
	"github.com/watarun54/serverless-skill-manager/server/interfaces/database"
	"github.com/watarun54/serverless-skill-manager/server/usecase"
)

type BillController struct {
	Interactor usecase.BillInteractor
}

func NewBillController(sqlHandler database.SqlHandler) *BillController {
	return &BillController{
		Interactor: usecase.BillInteractor{
			BillRepository: &database.BillRepository{
				SqlHandler: sqlHandler,
			},
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *BillController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	com := domain.Bill{
		ID: id,
	}
	bill, err := controller.Interactor.Bill(com)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(bill))
	return
}

func (controller *BillController) Create(c Context) (err error) {
	bForm := domain.BillForm{}
	c.Bind(&bForm)
	b, err := controller.Interactor.ConvertBillFormToBill(bForm)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	bill, err := controller.Interactor.Add(b)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, bill)
	return
}

func (controller *BillController) Update(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	bForm := domain.BillForm{}
	bForm.ID = id
	c.Bind(&bForm)
	b, err := controller.Interactor.ConvertBillFormToBill(bForm)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	bill, err := controller.Interactor.Update(b)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, NewResponse(bill))
	return
}

func (controller *BillController) Delete(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	bill := domain.Bill{
		ID: id,
	}
	err = controller.Interactor.DeleteById(bill)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, NewResponse(bill))
	return
}
