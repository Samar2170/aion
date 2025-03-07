package nasa

import (
	"aion/config"
	"aion/pkg/client"
	"encoding/json"
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

func (nc NasaClient) FetchAstronomyPhotoOfTheDay() (AstronomyPhotoOfTheDayResponse, error) {
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

	response, err := http.DefaultClient.Do(&request)
	if err != nil {
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
