package initialization

import (
	"github.com/glebarez/sqlite"
	"github.com/shaco-go/sqltostruct/internal/model"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.Option{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
