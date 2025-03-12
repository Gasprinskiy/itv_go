package repository

import (
	"itv_go/internal/entity/movie"
	appuser "itv_go/internal/entity/user"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(tx *gorm.DB, param appuser.CreateUserParams) (int, error)
	GetUserByLogin(tx *gorm.DB, login string) (appuser.User, error)
}

type Movie interface {
	CreateMovieRecord(tx *gorm.DB, param movie.CreateMovieRecordParam) (int, error)
	GetMovieByID(tx *gorm.DB, id int) (movie.Movie, error)
	UpdateMovie(tx *gorm.DB, param movie.Movie) (int, error)
	DeleteMovie(tx *gorm.DB, id int) (int, error)
	GetMovieList(tx *gorm.DB) ([]movie.Movie, error)
}
