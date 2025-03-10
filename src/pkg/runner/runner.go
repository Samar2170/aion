package runner

import (
	"aion/pkg/logging"
)

type FunctionWithParams struct {
	Name   string
	Fn     func(interface{}) error
	Params []interface{}
}

func RunJob(fwp FunctionWithParams) {
	logging.AuditLogger.Info().Msg("Running job " + fwp.Name)
	FN := fwp.Fn
	err := FN(fwp.Params)
	if err != nil {
		logging.ErrorLogger.Error().Msg("error while running job " + fwp.Name + " " + err.Error())
	}

}
