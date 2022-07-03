package domain

import "time"

type Rooms []Room

type (
	Room struct {
		ID          int          `json:"-"`
		UUID        string       `json:"id"`
		Name        string       `json:"name"`
		Users       []User       `json:"users" gorm:"many2many:user_rooms;save_associations:false;"`
		Bills       []Bill       `json:"-"`
		RoomMembers []RoomMember `json:"members"`
		CreatedAt   time.Time    `json:"created_at"`
		UpdatedAt   time.Time    `json:"updated_at"`
	}

	RoomForm struct {
		Room
		UserIds []int `json:"user_ids"`
	}
)
