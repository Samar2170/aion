package utils

import "time"

func CheckDateFormat(datestring, format string) bool {
	_, err := time.Parse(format, datestring)
	return err == nil
}
