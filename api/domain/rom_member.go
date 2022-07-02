package domain

type RoomMembers []RoomMember

type (
	RoomMember struct {
		ID     int     `json:"id"`
		Name   *string `json:"name" gorm:"unique;not null;size:255"`
		RoomID int     `json:"-" gorm:"not null"`
	}
)
