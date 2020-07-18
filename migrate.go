package main

import (
	"Project/migrate/config"
	"Project/migrate/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	conf := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.Username, config.DBConfig.Database, config.DBConfig.Password, config.DBConfig.SSLMode,
	)

	db, err := gorm.Open("postgres", conf)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	migrate(db)
}

func migrate(db *gorm.DB) {
	models := []interface{}{
		&model.Brand{},
		&model.Product{},
		&model.Purchase{},
	}
	db.AutoMigrate(models...)
}
