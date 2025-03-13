package postgress

import (
	appuser "itv_go/internal/entity/user"
	"itv_go/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() repository.User {
	return &userRepository{}
}

func (r *userRepository) CreateUser(tx *gorm.DB, param appuser.CreateUserParams) (int, error) {
	newRecord := appuser.User{}

	result := tx.Model(&appuser.User{}).Create(map[string]interface{}{
		"login":    param.Login,
		"password": param.Password,
	}).Last(&newRecord)

	return newRecord.ID, result.Error
}

func (r *userRepository) GetUserByLogin(tx *gorm.DB, login string) (appuser.User, error) {
	record := appuser.User{}

	result := tx.Find(&record, "login = ?", login)

	return record, result.Error
}
