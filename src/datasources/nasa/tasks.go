package nasa

import (
	"aion/pkg/client"
	"aion/pkg/db"
	"aion/pkg/logging"
	"aion/pkg/utils"
	"errors"
	"strings"
)

func GetAstronomyPhotoOfTheDay(dateString string) error {
	// dateString, ok := date.(string)
	// if !ok {
	// 	return errors.New("invalid date format")
	// }
	checkDateFormat := utils.CheckDateFormat(dateString, "2006-01-02")
	if !checkDateFormat {
		return errors.New("invalid date format")
	}
	nasaClient := NewNasaClient()
	resp, err := nasaClient.FetchAstronomyPhotoOfTheDay(dateString)
	if err != nil {
		return utils.HandleError(err, "Nasa")
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
		return utils.HandleError(err, "Nasa")
	}
	fu := client.NewFileupClient()
	image, err := nasaClient.Client.DownloadFile(resp.Hdurl, map[string]string{
		"api_key": nasaClient.Client.ApiKey,
	})
	if err != nil {
		return errors.New("error while downloading image " + err.Error())
	}
	urlSplit := strings.Split(resp.Hdurl, "/")
	filename := urlSplit[len(urlSplit)-1]
	err = fu.UploadFile(&image, filename)
	if err != nil {
		logging.ErrorLogger.Println("error while uploading image " + err.Error())
		return err
	}
	return nil
}
