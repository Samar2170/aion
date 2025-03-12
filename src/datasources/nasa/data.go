package nasa

import "aion/pkg/db"

func GetNasaApod() ([]NasaPhoto, error) {
	var nasaPhotos []NasaPhoto
	err := db.DB.Find(&nasaPhotos).Error
	if err != nil {
		return nil, err
	}
	return nasaPhotos, nil
}
