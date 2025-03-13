package usecase

import (
	"fmt"
	"itv_go/internal/entity/movie"
	"itv_go/internal/repository"
	transactiongeneric "itv_go/tools/transaction-generic"

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
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *gorm.DB) (int, error) {
			return u.movieRepo.CreateMovieRecord(tx, param)
		},
	)
}

func (u *MovieUsecase) GetMovieByID(id int) (movie.Movie, error) {
	fmt.Println("FUCKING MOVIE")
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *gorm.DB) (movie.Movie, error) {
			return u.movieRepo.GetMovieByID(tx, id)
		},
	)
}

func (u *MovieUsecase) UpdateMovie(param movie.Movie) (int, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *gorm.DB) (int, error) {
			return u.movieRepo.UpdateMovie(tx, param)
		},
	)
}

func (u *MovieUsecase) DeleteMovie(id int) (int, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *gorm.DB) (int, error) {
			return u.movieRepo.DeleteMovie(tx, id)
		},
	)
}

func (u *MovieUsecase) GetMovieList() ([]movie.Movie, error) {
	return transactiongeneric.HandleMethodWithTransaction(
		u.db,
		func(tx *gorm.DB) ([]movie.Movie, error) {
			return u.movieRepo.GetMovieList(tx)
		},
	)
}
