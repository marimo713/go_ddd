package mysql

import (
	"fmt"
	"log"
	"my-app/config"
	"my-app/infrastructure/mysql/gorm_model"

	"github.com/jinzhu/gorm"
)

type Database interface {
	conn() *gorm.DB
	autoMigrate() error
}

type database struct {
	db *gorm.DB
}

func Open(conf config.DBConfig) (Database, func(), error) {
	dbURL := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo", conf.User, conf.Password, conf.Host, conf.Port, conf.Name)

	gDB, err := gorm.Open("mysql", dbURL)
	if err != nil {
		return nil, nil, err
	}

	// DB切断用の関数を定義
	cleanup := func() {
		if err := gDB.Close(); err != nil {
			log.Print(err)
		}
	}

	db := database{
		db: gDB,
	}
	if err := db.autoMigrate(); err != nil {
		cleanup()
		return nil, nil, err
	}

	return &db, cleanup, nil

}

func (database *database) conn() *gorm.DB {
	return database.db
}

func (database *database) autoMigrate() error {
	return database.db.AutoMigrate(&gorm_model.Book{}).Error
}
