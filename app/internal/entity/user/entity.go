package appuser

type User struct {
	ID       int    `json:"id" db:"id" gorm:"primaryKey"`
	Login    string `json:"login" db:"login" gorm:"size:255;not null"`
	Password string `json:"password" db:"password" gorm:"size:255;not null"`
}

type CreateUserParams struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type UserResponse struct {
	ID int `json:"id"`
}

func NewUserFromCreateUserParams(param CreateUserParams) User {
	return User{
		ID:       -1,
		Login:    param.Login,
		Password: param.Password,
	}
}
