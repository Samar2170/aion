package client

import (
	"aion/pkg/db"
	"aion/pkg/logging"
	"bytes"
	"io"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func init() {
	db.DB.AutoMigrate(&RequestLog{})
}

type Client struct {
	Name    string
	BaseUrl string
	ApiKey  string
}

type RequestLog struct {
	*gorm.Model
	Url        string
	StatusCode int
	Time       time.Time
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultClient.Do(request)
	rq := RequestLog{
		Url:        request.URL.String(),
		StatusCode: response.StatusCode,
		Time:       time.Now(),
	}
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return nil, err
	}
	err = db.DB.Create(&rq).Error
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
	}
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return nil, err
	}
	return response, nil
}

func (c *Client) DownloadFile(url string, headers map[string]string) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	request, err := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return buffer, err
	}
	response, err := c.Do(request)
	if err != nil {
		logging.ErrorLogger.Error().Err(err)
		return buffer, err
	}
	defer response.Body.Close()
	if _, err := io.Copy(&buffer, response.Body); err != nil {
		logging.ErrorLogger.Error().Err(err)
		return buffer, err
	}
	return buffer, nil
}
