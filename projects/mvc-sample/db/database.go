package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=maodou24 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	pg, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	db = pg
}

func GetDb() *gorm.DB {
	return db
}
