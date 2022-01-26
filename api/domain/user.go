package domain

type Users []User

type (
	User struct {
		ID             int    `json:"id"`
		Name           string `json:"name" gorm:"not null;size:255"`
		Email          string `json:"email" gorm:"unique;not null;size:255"`
		HashedPassword string `json:"-"`
		LineID         string `json:"line_id"`
		Papers         []Paper
	}

	UserForm struct {
		User
		Password string `json:"password"`
	}
)
