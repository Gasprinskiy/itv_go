package repository

import (
	"database/sql"
	"itv_go/internal/entity/movie"
)

type MovieRepository interface {
	CreateMovieRecord(tx *sql.Tx, param movie.CreateMovieRecordParam) (int, error)
	GetMovieByID(tx *sql.Tx, id int) (movie.Movie, error)
	UpdateMovie(tx *sql.Tx, param movie.Movie) (int, error)
	DeleteMovie(tx *sql.Tx, id int) (int, error)
	GetMovieList(tx *sql.Tx, id int) ([]movie.Movie, error)
}
