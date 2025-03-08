package client

import (
	"aion/config"
	"aion/pkg/logging"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
)

type FileupClient struct {
	Client   Client
	Username string
}

func NewFileupClient() *FileupClient {
	return &FileupClient{
		Client: Client{
			Name:    "Fileup",
			BaseUrl: "https://assets.barebasics.shop/",
			ApiKey:  config.FileupAPIKey,
		},
		Username: config.FileupUsername,
	}
}

func (fc *FileupClient) Name() string {
	return "Fileup"
}

func (fc *FileupClient) UploadFile(r io.Reader, fileName string) error {
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return err
	}
	_, err = io.Copy(part, r)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return err
	}
	err = writer.Close()
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return err
	}
	req, err := http.NewRequest("POST", fc.Client.BaseUrl+"files/upload/", &body)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return err
	}
	req.Header.Set("X-API-Key", fc.Client.ApiKey)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return err
	}
	if response.StatusCode != 200 {
		errorMsg, err := io.ReadAll(response.Body)
		if err != nil {
			logging.ErrorLogger.Error().Err(err)
			return errors.New("failed to upload file " + response.Status)
		}
		return errors.New("failed to upload file " + response.Status + " " + string(errorMsg))
	}
	return nil
}

type FileSignedUrlResponse struct {
	Message string `json:"message"`
}

func (fc *FileupClient) GetFileUrl(fileName string) (string, error) {
	url := fc.Client.BaseUrl + "files/get-signed-url/" + fc.Username + fileName
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return "", err
	}
	request.Header.Set("X-API-Key", fc.Client.ApiKey)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		logging.ErrorLogger.Error().Msg("failed to get file url " + response.Status)
		errorMsg, err := io.ReadAll(response.Body)
		if err != nil {
			logging.ErrorLogger.Error().Err(err)
			return "", errors.New("failed to get file url " + response.Status)
		}
		return "", errors.New("failed to get file url " + response.Status + " " + string(errorMsg))
	}
	var result FileSignedUrlResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return "", err
	}
	return result.Message, nil
}
