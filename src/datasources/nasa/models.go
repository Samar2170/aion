package nasa

import (
	"aion/pkg/db"

	"gorm.io/gorm"
)

func init() {
	db.DB.AutoMigrate(&NasaPhoto{})
}

type NasaPhoto struct {
	*gorm.Model
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	Url         string `json:"url"`
	Date        string `json:"date"`
	CopyRight   string `json:"copyright"`

	SavedFileLink string `json:"saved_file_link"`
}
