package nasa

import (
	"aion/config"
	"aion/pkg/client"
	"aion/pkg/logging"
	"aion/pkg/utils"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path/filepath"
)

type NasaClient struct {
	Client    client.Client
	PhotosDir string
}

func NewNasaClient() *NasaClient {
	return &NasaClient{
		Client: client.Client{
			Name:    "Nasa",
			BaseUrl: "https://images-api.nasa.gov/",
			ApiKey:  config.NasaAPIKey,
		},
		PhotosDir: filepath.Join(
			config.BaseDir, config.BasePhotosDir,
			"nasa",
		),
	}
}

type AstronomyPhotoOfTheDayResponse struct {
	Copyright   string `json:"copyright"`
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	MediaType   string `json:"media_type"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

func (nc NasaClient) FetchAstronomyPhotoOfTheDay(dateString string) (AstronomyPhotoOfTheDayResponse, error) {
	checkDateFormat := utils.CheckDateFormat(dateString, "2006-01-02")
	if !checkDateFormat {
		return AstronomyPhotoOfTheDayResponse{}, errors.New("invalid date format")
	}
	subUrl := "planetary/apod"
	request := http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   "api.nasa.gov",
			Path:   subUrl,
		},
		Header: http.Header{
			"api_key": {nc.Client.ApiKey},
		},
	}

	response, err := nc.Client.Do(&request)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return AstronomyPhotoOfTheDayResponse{}, err
	}

	defer response.Body.Close()
	var result AstronomyPhotoOfTheDayResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return AstronomyPhotoOfTheDayResponse{}, err
	}
	return result, nil
}
