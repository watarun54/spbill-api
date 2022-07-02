package database

import (
	"github.com/watarun54/spbill-api/server/domain"
)

type RoomRepository struct {
	SqlHandler
}

func (repo *RoomRepository) FindOne(r domain.Room) (room domain.Room, err error) {
	if err = repo.Debug().Take(&room, r.ID).
		Related(&room.Users, "Users").
		Related(&room.Bills, "Bills").
		Related(&room.RoomMembers, "RoomMembers").
		Error; err != nil {
		return
	}
	return
}

func (repo *RoomRepository) FindAll() (rooms domain.Rooms, err error) {
	if err = repo.Debug().Preload("Users").Find(&rooms).Error; err != nil {
		return
	}
	return
}

func (repo *RoomRepository) Store(r domain.Room) (room domain.Room, err error) {
	if err = repo.Debug().Create(&r).Error; err != nil {
		return
	}
	if err = repo.Debug().Model(&r).Association("Users").Replace(r.Users).Error; err != nil {
		return
	}
	room = r
	return
}

func (repo *RoomRepository) Update(r domain.Room) (room domain.Room, err error) {
	if err = repo.Debug().Set("gorm:save_associations", false).Take(&domain.Room{ID: r.ID}).Updates(&r).Error; err != nil {
		return
	}
	if err = repo.Debug().Model(&r).Association("Users").Replace(r.Users).Error; err != nil {
		return
	}
	room = r
	return
}

func (repo *RoomRepository) DeleteById(r domain.Room) (err error) {
	if err = repo.Debug().Model(&r).Association("Users").Replace(&[]domain.User{}).Error; err != nil {
		return
	}
	if err = repo.Debug().Delete(&r).Error; err != nil {
		return
	}
	return
}
