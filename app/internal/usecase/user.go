package usecase

import (
	"itv_go/internal/entity/global"
	appuser "itv_go/internal/entity/user"
	"itv_go/internal/repository"
	"itv_go/tools/passencoder"
	transactiongeneric "itv_go/tools/transaction-generic"

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
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	defer tx.Commit()

	tx.Begin()

	_, err := u.userRepo.GetUserByLogin(tx, param.Login)

	switch err {
	case nil:
		tx.Rollback()
		return 0, global.ErrUserAllreadyExists

	case gorm.ErrRecordNotFound:
		encodePass, err := passencoder.CreateHashPassword(param.Password)
		if err != nil {
			return 0, global.ErrInternalError
		}

		param.Password = encodePass

		id, err := u.userRepo.CreateUser(tx, param)
		if err != nil {
			tx.Rollback()
			err = global.ErrInternalError
		}

		return id, err

	default:
		tx.Rollback()
		return 0, global.ErrInternalError
	}
}

func (u *UserUsecase) Auth(param appuser.CreateUserParams) (int, error) {
	user, err := transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *gorm.DB) (appuser.User, error) {
			return u.userRepo.GetUserByLogin(tx, param.Login)
		},
	)

	if err != nil {
		return 0, err
	}

	isPassCorrect := passencoder.CheckHashPassword(user.Password, param.Password)
	if !isPassCorrect {
		return 0, global.ErrInvalidLoginOrPassword
	}

	return user.ID, nil
}
