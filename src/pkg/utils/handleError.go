package utils

import "aion/pkg/logging"

func HandleError(err error, handler string) error {
	logging.Errorlogger.Error().Err(err)

	return err
}
