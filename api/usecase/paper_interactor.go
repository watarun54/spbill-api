package usecase

import "github.com/watarun54/serverless-skill-manager/server/domain"

type (
	IPaperRepository interface {
		FindOne(c domain.Paper) (domain.Paper, error)
		FindAll(c domain.Paper) (domain.Papers, error)
		Store(c domain.Paper) (domain.Paper, error)
		Update(c domain.Paper) (domain.Paper, error)
		DeleteById(c domain.Paper) error
	}

	PaperInteractor struct {
		PaperRepository IPaperRepository
	}
)

func (interactor *PaperInteractor) Paper(c domain.Paper) (paper domain.Paper, err error) {
	paper, err = interactor.PaperRepository.FindOne(c)
	return
}

func (interactor *PaperInteractor) Papers(c domain.Paper) (papers domain.Papers, err error) {
	papers, err = interactor.PaperRepository.FindAll(c)
	return
}

func (interactor *PaperInteractor) Add(c domain.Paper) (paper domain.Paper, err error) {
	paper, err = interactor.PaperRepository.Store(c)
	return
}

func (interactor *PaperInteractor) Update(c domain.Paper) (paper domain.Paper, err error) {
	paper, err = interactor.PaperRepository.Update(c)
	return
}

func (interactor *PaperInteractor) DeleteById(c domain.Paper) (err error) {
	err = interactor.PaperRepository.DeleteById(c)
	return
}
