package database

import (
	"github.com/watarun54/spbill-api/server/domain"
)

type RoomMemberRepository struct {
	SqlHandler
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
