package nasa

import (
	"aion/pkg/db"
	"aion/pkg/utils"

	"gorm.io/gorm"
)

type NasaPhoto struct {
	*gorm.Model
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	Url         string `json:"url"`
	Date        string `json:"date"`
	CopyRight   string `json:"copyright"`
}
type NasaHandler struct {
}

func (nh *NasaHandler) Name() string {
	return "Nasa"
}

func (nh *NasaHandler) GetAstronomyPhotoOfTheDay() error {
	nasaClient := NewNasaClient()
	resp, err := nasaClient.FetchAstronomyPhotoOfTheDay()
	if err != nil {
		return utils.HandleError(err, nh.Name())
	}
	np := NasaPhoto{
		Title:       resp.Title,
		Explanation: resp.Explanation,
		Hdurl:       resp.Hdurl,
		Url:         resp.Url,
		Date:        resp.Date,
		CopyRight:   resp.Copyright,
	}
	err = db.DB.Create(&np).Error
	if err != nil {
		return utils.HandleError(err, nh.Name())
	}
	return nil
}
