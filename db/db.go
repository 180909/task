package db

import (
	"context"

	"github.com/180909/task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	dsn := "host=localhost user=wang password=root dbname=corn port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Job{})
	DB = db
}

func GetAllJobs(ctx context.Context) ([]models.Job, error) {
	var jobs []models.Job
	tx := DB.Find(&jobs)
	return jobs, tx.Error
}
