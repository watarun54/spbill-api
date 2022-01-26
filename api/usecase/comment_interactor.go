package usecase

import "github.com/watarun54/serverless-skill-manager/server/domain"

type (
	ICommentRepository interface {
		FindOne(c domain.Comment) (domain.Comment, error)
		FindAll(c domain.Comment) (domain.Comments, error)
		Store(c domain.Comment) (domain.Comment, error)
		Update(c domain.Comment) (domain.Comment, error)
		DeleteById(c domain.Comment) error
	}

	CommentInteractor struct {
		CommentRepository ICommentRepository
	}
)

func (interactor *CommentInteractor) Comment(c domain.Comment) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRepository.FindOne(c)
	return
}

func (interactor *CommentInteractor) Comments(c domain.Comment) (comments domain.Comments, err error) {
	comments, err = interactor.CommentRepository.FindAll(c)
	return
}

func (interactor *CommentInteractor) Add(c domain.Comment) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRepository.Store(c)
	return
}

func (interactor *CommentInteractor) Update(c domain.Comment) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRepository.Update(c)
	return
}

func (interactor *CommentInteractor) DeleteById(c domain.Comment) (err error) {
	err = interactor.CommentRepository.DeleteById(c)
	return
}
