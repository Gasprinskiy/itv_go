package repository

import (
	"itv_go/internal/entity/movie"

	"gorm.io/gorm"
)

type Movie interface {
	CreateMovieRecord(tx *gorm.DB, param movie.CreateMovieRecordParam) (int, error)
	GetMovieByID(tx *gorm.DB, id int) (movie.Movie, error)
	UpdateMovie(tx *gorm.DB, param movie.Movie) (int, error)
	DeleteMovie(tx *gorm.DB, id int) (int, error)
	GetMovieList(tx *gorm.DB) ([]movie.Movie, error)
}
