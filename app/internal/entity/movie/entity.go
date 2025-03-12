package movie

import "time"

type Movie struct {
	ID       int       `json:"id" db:"id" gorm:"primaryKey"`
	Title    string    `json:"title" db:"title" gorm:"size:255;not null"`
	Director string    `json:"director" db:"director" gorm:"size:255;not null"`
	Plot     string    `json:"plot" db:"plot" gorm:"size:255;not null"`
	Year     time.Time `json:"year" db:"year" gorm:"type:date;not null"`
}

type CreateMovieRecordParam struct {
	Title    string    `json:"title" db:"title"`
	Director string    `json:"director" db:"director"`
	Plot     string    `json:"plot" db:"plot"`
	Year     time.Time `json:"year" db:"year"`
}
