package database

import (
	"github.com/watarun54/serverless-skill-manager/server/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	if err = repo.Debug().Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByIds(ids []int) (users domain.Users, err error) {
	if err = repo.Debug().Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByEmail(email string) (user domain.User, err error) {
	if err = repo.Debug().Where("email = ?", email).Find(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindByLineID(lineID string) (user domain.User, err error) {
	if err = repo.Debug().Where("line_id = ?", lineID).Take(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	if err = repo.Debug().Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = repo.Debug().Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = repo.Debug().Take(&domain.User{}, u.ID).Updates(&u).Scan(&user).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) DeleteById(user domain.User) (err error) {
	if err = repo.Debug().Where("user_id = ?", user.ID).Delete(user.Papers).Error; err != nil {
		return
	}
	if err = repo.Debug().Delete(&user).Error; err != nil {
		return
	}
	return
}
