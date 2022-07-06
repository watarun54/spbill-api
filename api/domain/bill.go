package domain

import "time"

type Bills []Bill

type (
	Bill struct {
		ID        int          `json:"id"`
		Name      string       `json:"name" gorm:"not null;size:255;"`
		Amount    int          `json:"amount" gorm:"not null;"`
		RoomID    int          `json:"-" gorm:"not null"`
		PayerID   int          `json:"-" gorm:"not null"`
		Payer     RoomMember   `json:"payer"`
		Payees    []RoomMember `json:"payees" gorm:"many2many:bill_payees;association_jointable_foreignkey:payee_id;"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
	}

	BillForm struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Amount   int    `json:"amount"`
		RoomID   int    `json:"room_id"`
		PayerID  int    `json:"payer_id"`
		PayeeIds []int  `json:"payee_ids"`
	}

	UserPaymentRes struct {
		FromUser User `json:"from_user"`
		ToUser   User `json:"to_user"`
		Amount   int  `json:"amount"`
	}
)
