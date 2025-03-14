package appuser

type User struct {
	ID       int    `json:"id" db:"id" gorm:"primaryKey"`
	Login    string `json:"login" db:"login" gorm:"size:255;not null"`
	Password string `json:"password" db:"password" gorm:"size:255;not null"`
}

type CreateUserParams struct {
	Login    string `json:"login" db:"login" validate:"required,min=4,max=32"`
	Password string `json:"password" db:"password" validate:"required,min=5,max=32"`
}
