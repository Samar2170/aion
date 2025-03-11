package nasa

import (
	"aion/config"
	"aion/pkg/client"
	"aion/pkg/logging"
	"encoding/json"
	"errors"
	"net/http"
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
			BaseUrl: "https://api.nasa.gov/",
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
	subUrl := "planetary/apod"
	qs := "date=" + dateString

	request, err := http.NewRequest("GET", nc.Client.BaseUrl+subUrl+"?"+qs, nil)
	if err != nil {
		logging.ErrorLogger.Println(err)
		return AstronomyPhotoOfTheDayResponse{}, err
	}
	request.Header.Set("X-API-Key", nc.Client.ApiKey)
	response, err := nc.Client.Do(request, "nasa")
	if err != nil {
		logging.ErrorLogger.Println(err)
		return AstronomyPhotoOfTheDayResponse{}, err
	}
	if response.StatusCode != 200 {
		return AstronomyPhotoOfTheDayResponse{}, errors.New("failed to fetch astronomy photo of the day response -" + response.Status)
	}

	defer response.Body.Close()
	var result AstronomyPhotoOfTheDayResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return AstronomyPhotoOfTheDayResponse{}, err
	}
	return result, nil
}
