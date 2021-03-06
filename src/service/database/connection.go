package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"merger/src/handler"
	"os"
	"sync"
)

type singleton struct {
	db *gorm.DB
}

var instance *singleton
var once sync.Once

func GetConnection() *gorm.DB {
	once.Do(func() {
		instance = &singleton{db: connect()}
	})

	return instance.db
}

func connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open("postgres", dsn)

	db.DB().SetMaxIdleConns(0)
	db.LogMode(false)

	handler.FailOnError(err, "Can't connect to database")

	return db
}
