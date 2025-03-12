package database

import (
	"fmt"
	"itv_go/config"
	"itv_go/internal/entity/movie"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(conf *config.Config) *gorm.DB {
	fmt.Println("FUCK: ", conf.DbUrl)
	db, err := gorm.Open(postgres.Open(conf.DbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic("Could not connect to database: ", err)
	}

	err = db.AutoMigrate(&movie.Movie{})
	if err != nil {
		log.Panic("Could not automigrate: ", err)
	}

	return db
}
