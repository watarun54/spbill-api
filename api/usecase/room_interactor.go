package usecase

import (
	"github.com/watarun54/spbill-api/server/domain"
)

type (
	IRoomRepository interface {
		FindOne(r domain.Room) (domain.Room, error)
		FindByUUID(uid string) (domain.Room, error)
		FindAll() (domain.Rooms, error)
		Store(r domain.Room) (domain.Room, error)
		Update(r domain.Room) (domain.Room, error)
		DeleteByUUID(r domain.Room) error
	}

	RoomInteractor struct {
		RoomRepository IRoomRepository
		UserRepository IUserRepository
	}
)

func (interactor *RoomInteractor) ConvertRoomFormToRoom(roomForm domain.RoomForm) (room domain.Room, err error) {
	users, err := interactor.UserRepository.FindByIds(roomForm.UserIds)
	if err != nil {
		return
	}
	room.UUID = roomForm.UUID
	room.Name = roomForm.Name
	room.Users = users
	return
}

func (interactor *RoomInteractor) Room(r domain.Room) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.FindOne(r)
	return
}

func (interactor *RoomInteractor) FindByUUID(uid string) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.FindByUUID(uid)
	return
}

func (interactor *RoomInteractor) Rooms() (rooms domain.Rooms, err error) {
	rooms, err = interactor.RoomRepository.FindAll()
	return
}

func (interactor *RoomInteractor) Add(r domain.Room) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.Store(r)
	return
}

func (interactor *RoomInteractor) Update(r domain.Room) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.Update(r)
	return
}

func (interactor *RoomInteractor) DeleteByUUID(r domain.Room) (err error) {
	err = interactor.RoomRepository.DeleteByUUID(r)
	return
}
