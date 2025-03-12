package postgress

import (
	"itv_go/internal/entity/movie"
	"itv_go/internal/repository"

	"gorm.io/gorm"
)

type movieRepository struct{}

func NewMovieRepository() repository.Movie {
	return &movieRepository{}
}

func (r *movieRepository) CreateMovieRecord(tx *gorm.DB, param movie.CreateMovieRecordParam) (int, error) {
	newRecord := movie.NewMovieFromCreateMovieRecordParam(param)

	result := tx.Create(&newRecord)

	return newRecord.ID, result.Error
}

func (r *movieRepository) GetMovieByID(tx *gorm.DB, id int) (movie.Movie, error) {
	record := movie.Movie{}

	result := tx.First(&record, id)

	return record, result.Error
}

func (r *movieRepository) UpdateMovie(tx *gorm.DB, param movie.Movie) (int, error) {
	result := tx.Save(&param)

	return param.ID, result.Error
}

func (r *movieRepository) DeleteMovie(tx *gorm.DB, id int) (int, error) {
	result := tx.Model(&movie.Movie{}).Where("id = ?", id).Update("deleted", true)

	return id, result.Error
}
func (r *movieRepository) GetMovieList(tx *gorm.DB) ([]movie.Movie, error) {
	listOfRecords := []movie.Movie{}

	result := tx.Find(&listOfRecords)

	return listOfRecords, result.Error
}
