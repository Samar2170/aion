package main

import (
	"aion/datasources/nasa"
	"aion/pkg/runner"
)

func main() {
	runner.RunJob(runner.FunctionWithParams{
		Name: "GetAstronomyPhotoOfTheDay",
		Fn:   nasa.GetAstronomyPhotoOfTheDay,
		Params: []interface{}{
			"2022-01-01",
		},
	})
}
