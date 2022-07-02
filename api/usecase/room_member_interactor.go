package usecase

import (
	"github.com/watarun54/spbill-api/server/domain"
)

type (
	IRoomMemberRepository interface {
		Store(r domain.RoomMember) (domain.RoomMember, error)
	}

	RoomMemberInteractor struct {
		RoomMemberRepository IRoomMemberRepository
	}
)

func (interactor *RoomMemberInteractor) Add(r domain.RoomMember) (roomMember domain.RoomMember, err error) {
	roomMember, err = interactor.RoomMemberRepository.Store(r)
	return
}
