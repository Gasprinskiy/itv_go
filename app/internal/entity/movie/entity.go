package movie

import "time"

type Movie struct {
	ID       int       `json:"id" db:"id" gorm:"primaryKey"`
	Title    string    `json:"title" db:"title" gorm:"size:255;not null"`
	Director string    `json:"director" db:"director" gorm:"size:255;not null"`
	Plot     string    `json:"plot" db:"plot" gorm:"size:255;not null"`
	Year     time.Time `json:"year" db:"year" gorm:"type:date;not null"`
	Deleted  bool      `db:"deleted" gorm:"default:false;not null"`
}

type CreateMovieRecordParam struct {
	Title    string    `json:"title" db:"title" validate:"required,max=50"`
	Director string    `json:"director" db:"director" validate:"required,max=25"`
	Plot     string    `json:"plot" db:"plot" validate:"required,min=25,max=250"`
	Year     time.Time `json:"year" db:"year" validate:"required"`
}

func NewMovieFromCreateMovieRecordParam(param CreateMovieRecordParam) Movie {
	return Movie{
		ID:       -1,
		Title:    param.Title,
		Director: param.Director,
		Plot:     param.Plot,
		Year:     param.Year,
	}
}
