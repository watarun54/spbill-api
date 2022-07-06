package usecase

import (
	"errors"
	"github.com/watarun54/spbill-api/server/domain"
)

type (
	IRoomMemberRepository interface {
		FindById(id int) (domain.RoomMember, error)
		FindByIds(ids []int) ([]domain.RoomMember, error)
		Store(r domain.RoomMember) (domain.RoomMember, error)
		Delete(domain.RoomMember) error
	}

	RoomMemberInteractor struct {
		RoomMemberRepository IRoomMemberRepository
		BillRepository       IBillRepository
	}
)

func (interactor *RoomMemberInteractor) Add(r domain.RoomMember) (roomMember domain.RoomMember, err error) {
	roomMember, err = interactor.RoomMemberRepository.Store(r)
	return
}

func (interactor *RoomMemberInteractor) Delete(member domain.RoomMember) (err error) {
	currentMember, err := interactor.RoomMemberRepository.FindById(member.ID)
	bill := domain.Bill{
		RoomID: currentMember.RoomID,
	}
	bills, err := interactor.BillRepository.FindAll(bill)
	if err != nil {
		return
	}
	for _, b := range bills {
		if b.Payer.ID == member.ID {
			err = errors.New("立替の支払った人として登録されているため、削除できません。")
			return
		}
		for _, payee := range b.Payees {
			if payee.ID == member.ID {
				err = errors.New("立替の支払ってもらった人として登録されているため、削除できません。")
				return
			}
		}
	}
	err = interactor.RoomMemberRepository.Delete(member)
	return
}
