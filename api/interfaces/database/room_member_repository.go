package database

import (
	"github.com/watarun54/spbill-api/server/domain"
)

type RoomMemberRepository struct {
	SqlHandler
}

func (repo *RoomMemberRepository) FindById(id int) (member domain.RoomMember, err error) {
	if err = repo.Debug().Find(&member, id).Error; err != nil {
		return
	}
	return
}

func (repo *RoomMemberRepository) FindByIds(ids []int) (members []domain.RoomMember, err error) {
	if err = repo.Debug().Where("id IN (?)", ids).Find(&members).Error; err != nil {
		return
	}
	return
}

func (repo *RoomMemberRepository) Store(rm domain.RoomMember) (roomMember domain.RoomMember, err error) {
	if err = repo.Debug().Create(&rm).Error; err != nil {
		return
	}
	if err = repo.Debug().Model(&rm).Error; err != nil {
		return
	}
	roomMember = rm
	return
}

func (repo *RoomMemberRepository) Delete(member domain.RoomMember) (err error) {
	if err = repo.Debug().Delete(&member).Error; err != nil {
		return
	}
	return
}
