package dbcontext

import (
	"Auction/domain/entity"
	"context"
	"fmt"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgContext struct {
	db *gorm.DB
}

func NewPgContext(lc fx.Lifecycle) *PgContext {
	dsn := "host=localhost user=postgres password=password dbname=auctiondb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Println("Closing database connection")
			return sqlDB.Close()
		},
	})

	return &PgContext{db: db}
}

func (ctx *PgContext) Context() *gorm.DB {
	return ctx.db
}

func (ctx PgContext) Migrate() {
	db := ctx.db
	db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Category{}, &entity.Lot{}, &entity.Photo{}, &entity.Bid{})
}
