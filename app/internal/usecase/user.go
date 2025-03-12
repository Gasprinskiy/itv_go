package usecase

import (
	"itv_go/internal/entity/global"
	appuser "itv_go/internal/entity/user"
	"itv_go/internal/repository"
	"itv_go/tools/passencoder"

	"gorm.io/gorm"
)

type UserUsecase struct {
	db       *gorm.DB
	jwt      *JwtUsecase
	userRepo repository.User
}

func NewUserUsecase(
	db *gorm.DB,
	jwt *JwtUsecase,
	userRepo repository.User,
) *UserUsecase {
	return &UserUsecase{
		db,
		jwt,
		userRepo,
	}
}

func (u *UserUsecase) Register(param appuser.CreateUserParams) (int, error) {
	encodePass, err := passencoder.CreateHashPassword(param.Password)
	if err != nil {
		return 0, global.ErrInternalError
	}

	param.Password = encodePass

	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	id, err := u.userRepo.CreateUser(tx, param)
	if err != nil {
		tx.Rollback()
		err = global.ErrInternalError
	}

	return id, err
}

func (u *UserUsecase) Auth(param appuser.CreateUserParams) (int, error) {
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	user, err := u.userRepo.GetUserByLogin(tx, param.Login)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, global.ErrInvalidLoginOrPassword
		}
		return 0, global.ErrInternalError
	}

	isPassCorrect := passencoder.CheckHashPassword(user.Password, param.Password)
	if !isPassCorrect {
		return 0, global.ErrInvalidLoginOrPassword
	}

	return user.ID, nil
}
