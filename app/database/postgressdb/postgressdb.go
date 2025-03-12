package postgressdb

// import (
// 	"gorm.io/gorm"
// )

// func NewDatabase() (*gorm.DB, error) {
// 	// dsn := "host=db user=user password=password dbname=mydb port=5432 sslmode=disable"
// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 	// 	Logger: logger.Default.LogMode(logger.Info), // Включаем логи SQL-запросов
// 	// })
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("failed to connect to database: %w", err)
// 	// }

// 	// // Настройка соединений (для продакшена)
// 	// sqlDB, err := db.DB()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// sqlDB.SetMaxIdleConns(10)
// 	// sqlDB.SetMaxOpenConns(100)
// 	// sqlDB.SetConnMaxLifetime(time.Hour)

// 	// log.Println("Connected to database!")
// 	// return db, nil
// }
