package cmd

import (
	"aion/pkg/db"
	"aion/pkg/logging"
	"errors"
	"time"
)

type FunctionWithParams struct {
	Name   string
	Fn     interface{}
	Params []interface{}
}

func RunJob(fwp FunctionWithParams) {
	logging.AuditLogger.Info().Msg("Running job " + fwp.Name)
	var err error
	// FN := fwp.Fn
	// err := FN(fwp.Params)
	switch fn := fwp.Fn.(type) {
	case func(string) error:
		if len(fwp.Params) > 0 {
			if str, ok := fwp.Params[0].(string); ok {
				err = fn(str)
			} else {
				err = errors.New("parameter is not a string")
			}
		} else {
			err = errors.New("no parameters provided")
		}
	case func(interface{}) error:
		err = fn(fwp.Params)
	}

	if err != nil {
		db.DB.Create(&Job{
			Name:     fwp.Name,
			RunAt:    time.Now(),
			HasError: true,
			ErrorMsg: err.Error(),
		})
		logging.ErrorLogger.Error().Msg("error while running job " + fwp.Name + " " + err.Error())
	} else {
		db.DB.Create(&Job{
			Name:     fwp.Name,
			RunAt:    time.Now(),
			HasError: false,
			ErrorMsg: "",
		})
	}
}

// func RunJob(job func() error) {
// 	jobName := reflect.ValueOf(job).String()

// 	err := job()
// 	if err != nil {
// 		db.DB.Create(&datasources.Job{
// 			Name:     jobName,
// 			RunAt:    time.Now(),
// 			HasError: true,
// 			ErrorMsg: err.Error(),
// 		})
// 	}
// }
