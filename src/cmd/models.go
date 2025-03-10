package cmd

import (
	"aion/pkg/db"
	"time"

	"gorm.io/gorm"
)

func init() {
	db.DB.AutoMigrate(&Job{})
}

type Job struct {
	*gorm.Model
	Name     string
	RunAt    time.Time
	HasError bool
	ErrorMsg string
}
