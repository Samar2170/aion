package datasources

import (
	"aion/pkg/db"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type Job struct {
	*gorm.Model
	Name     string
	RunAt    time.Time
	HasError bool
	ErrorMsg string
}

func RunJob(job func() error) {
	jobName := reflect.ValueOf(job).String()

	err := job()
	if err != nil {
		db.DB.Create(&Job{
			Name:     jobName,
			RunAt:    time.Now(),
			HasError: true,
			ErrorMsg: err.Error(),
		})
	}
}
