package dbcontext

import (
	"Auction/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func pgContext() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=auctiondb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}

func Migrate() {
	db, err := pgContext()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Category{}, &entity.Lot{}, &entity.Photo{}, &entity.Bid{})
}
