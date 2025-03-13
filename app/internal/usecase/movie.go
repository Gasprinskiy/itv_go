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

func (u *MovieUsecase) GetMovieByID(id int) (movie.Movie, error) {
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	defer tx.Commit()

	tx.Begin()

	movie, err := u.movieRepo.GetMovieByID(tx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = global.ErrNoData
		} else {
			err = global.ErrInternalError
		}
		tx.Rollback()
	}

	return movie, err
}

func (u *MovieUsecase) UpdateMovie(param movie.Movie) (int, error) {
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	defer tx.Commit()

	tx.Begin()

	id, err := u.movieRepo.UpdateMovie(tx, param)
	if err != nil {
		err = global.ErrInternalError
		tx.Rollback()
	}

	return id, err
}

func (u *MovieUsecase) DeleteMovie(id int) (int, error) {
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	defer tx.Commit()

	tx.Begin()

	id, err := u.movieRepo.DeleteMovie(tx, id)
	if err != nil {
		err = global.ErrInternalError
		tx.Rollback()
	}

	return id, err
}

func (u *MovieUsecase) GetMovieList() ([]movie.Movie, error) {
	tx := u.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
	defer tx.Commit()

	tx.Begin()

	list, err := u.movieRepo.GetMovieList(tx)
	if err != nil {
		err = global.ErrInternalError
		tx.Rollback()
	}

	return list, err
}
