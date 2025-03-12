package nasa

import (
	"aion/config"
	"aion/pkg/db"
	"aion/pkg/logging"
	"aion/pkg/utils"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetAstronomyPhotoOfTheDay(dateString string) error {
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
	// fu := client.NewFileupClient()
	image, err := nasaClient.Client.DownloadFile(resp.Hdurl, map[string]string{
		"api_key": nasaClient.Client.ApiKey,
	})
	if err != nil {
		return errors.New("error while downloading image " + err.Error())
	}
	urlSplit := strings.Split(resp.Hdurl, "/")
	filename := urlSplit[len(urlSplit)-1]
	file, err := os.Create(filepath.Join(config.BaseDir, config.BaseDataDir, NasaPhotosDir, filename))
	if err != nil {
		logging.ErrorLogger.Println("error while creating image locally" + err.Error())
		return err
	}
	defer file.Close()
	_, err = file.Write(image.Bytes())
	if err != nil {
		logging.ErrorLogger.Println("error while writing to dst file" + err.Error())
		return err
	}
	// err = fu.UploadFile(&image, filename)
	// if err != nil {
	// 	logging.ErrorLogger.Println("error while uploading image " + err.Error())
	// 	return err
	// }
	return nil
}

const NasaPhotosDir = "photos/nasa"
