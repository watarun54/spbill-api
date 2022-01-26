package usecase

import (
	"github.com/watarun54/serverless-skill-manager/server/domain"
)

type (
	IRoomRepository interface {
		FindOne(r domain.Room) (domain.Room, error)
		FindAll() (domain.Rooms, error)
		Store(r domain.Room) (domain.Room, error)
		Update(r domain.Room) (domain.Room, error)
		DeleteById(r domain.Room) error
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
	room.ID = roomForm.ID
	room.Name = roomForm.Name
	room.Users = users
	return
}

func (interactor *RoomInteractor) Room(r domain.Room) (room domain.Room, err error) {
	room, err = interactor.RoomRepository.FindOne(r)
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

func (interactor *RoomInteractor) DeleteById(r domain.Room) (err error) {
	err = interactor.RoomRepository.DeleteById(r)
	return
}
