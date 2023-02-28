package models

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Database(path string) *gorm.DB {
	var err error
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("error while connecting with mysql", err)
		return nil
	}

	//sql connect
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	err = db.AutoMigrate(&Customer{})

	if err != nil {
		log.Fatal("error while autoMigrate", err)
		return nil
	}

	return db
}
