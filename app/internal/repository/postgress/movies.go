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
	newRecord := movie.Movie{}

	result := tx.Model(&movie.Movie{}).Create(map[string]interface{}{
		"title":    param.Title,
		"director": param.Director,
		"plot":     param.Plot,
		"year":     param.Year,
	})
	tx.Last(&newRecord)

	return newRecord.ID, result.Error
}

func (r *movieRepository) GetMovieByID(tx *gorm.DB, id int) (movie.Movie, error) {
	record := movie.Movie{}

	result := tx.Model(&movie.Movie{}).Where("id = ? AND deleted = ?", id, false).First(&record)

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

	result := tx.Model(&movie.Movie{}).Where("deleted = ?", false).Find(&listOfRecords)

	return listOfRecords, result.Error
}
