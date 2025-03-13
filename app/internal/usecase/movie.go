package usecase

import (
	"itv_go/internal/entity/global"
	"itv_go/internal/entity/movie"
	"itv_go/internal/repository"

	"gorm.io/gorm"
)

type MovieUsecase struct {
	db        *gorm.DB
	movieRepo repository.Movie
}

func NewMovieUsecase(
	db *gorm.DB,
	movieRepo repository.Movie,
) *MovieUsecase {
	return &MovieUsecase{db, movieRepo}
}

func (u *MovieUsecase) CreateNewMovieRecord(param movie.CreateMovieRecordParam) (int, error) {
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	defer tx.Commit()

	tx.Begin()

	id, err := u.movieRepo.CreateMovieRecord(tx, param)
	if err != nil {
		err = global.ErrInternalError
		tx.Rollback()
	}

	return id, err
}
