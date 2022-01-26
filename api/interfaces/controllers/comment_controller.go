package controllers

import (
	"strconv"

	"github.com/watarun54/serverless-skill-manager/server/domain"
	"github.com/watarun54/serverless-skill-manager/server/interfaces/database"
	"github.com/watarun54/serverless-skill-manager/server/usecase"
)

type CommentController struct {
	Interactor usecase.CommentInteractor
}

func NewCommentController(sqlHandler database.SqlHandler) *CommentController {
	return &CommentController{
		Interactor: usecase.CommentInteractor{
			CommentRepository: &database.CommentRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CommentController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := userIDFromToken(c)
	com := domain.Comment{
		ID:     id,
		UserID: uid,
	}
	comment, err := controller.Interactor.Comment(com)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comment)
	return
}

func (controller *CommentController) Index(c Context) (err error) {
	uid := userIDFromToken(c)
	com := domain.Comment{
		UserID: uid,
	}
	comments, err := controller.Interactor.Comments(com)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comments)
	return
}

func (controller *CommentController) Create(c Context) (err error) {
	uid := userIDFromToken(c)
	com := domain.Comment{
		UserID: uid,
	}
	c.Bind(&com)
	comment, err := controller.Interactor.Add(com)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comment)
	return
}

func (controller *CommentController) Update(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := userIDFromToken(c)
	com := domain.Comment{
		ID:     id,
		UserID: uid,
	}
	c.Bind(&com)
	comment, err := controller.Interactor.Update(com)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comment)
	return
}

func (controller *CommentController) Delete(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := userIDFromToken(c)
	com := domain.Comment{
		ID:     id,
		UserID: uid,
	}
	err = controller.Interactor.DeleteById(com)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, com)
	return
}
