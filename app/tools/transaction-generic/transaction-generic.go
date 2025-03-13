package transactiongeneric

import (
	"itv_go/internal/entity/global"

	"gorm.io/gorm"
)

func HandleMethodWithTransaction[T any](
	db *gorm.DB,
	fn func(tx *gorm.DB) (T, error),
) (T, error) {
	var result T

	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	tx.Begin()

	result, err := fn(tx)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = global.ErrNoData
		} else {
			err = global.ErrInternalError
		}
		tx.Rollback()
	}

	return result, err
}
