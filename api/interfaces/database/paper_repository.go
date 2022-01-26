package database

import (
	"github.com/watarun54/serverless-skill-manager/server/domain"
)

type PaperRepository struct {
	SqlHandler
}

func (repo *PaperRepository) FindOne(p domain.Paper) (paper domain.Paper, err error) {
	if err = repo.Debug().Where(&p).Take(&paper).Error; err != nil {
		return
	}
	return
}

func (repo *PaperRepository) FindAll(p domain.Paper) (papers domain.Papers, err error) {
	if err = repo.Debug().Where(&p).Find(&papers).Error; err != nil {
		return
	}
	return
}

func (repo *PaperRepository) Store(p domain.Paper) (paper domain.Paper, err error) {
	if err = repo.Debug().Create(&p).Error; err != nil {
		return
	}
	paper = p
	return
}

func (repo *PaperRepository) Update(p domain.Paper) (paper domain.Paper, err error) {
	if err = repo.Debug().Take(&domain.Paper{ID: p.ID}, "user_id = ?", p.UserID).Updates(&p).Scan(&paper).Error; err != nil {
		return
	}
	return
}

func (repo *PaperRepository) DeleteById(p domain.Paper) (err error) {
	if err = repo.Debug().Take(&domain.Paper{ID: p.ID}, "user_id = ?", p.UserID).Delete(&p).Error; err != nil {
		return
	}
	return
}
