package config

import (
	"log"
	"voting-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {

	dsn := "root:senha@tcp(127.0.0.1:8893)/voting_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	db.AutoMigrate(&models.Candidate{}, &models.Votes{})
	return db
}
