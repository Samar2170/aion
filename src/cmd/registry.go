package cmd

import (
	"aion/datasources/nasa"
	"errors"
)

var registry = map[string]FunctionWithParams{
	"GetAstronomyPhotoOfTheDay": {
		Name: "GetAstronomyPhotoOfTheDay",
		Fn:   nasa.GetAstronomyPhotoOfTheDay,
		Params: []interface{}{
			"2025-01-01",
		},
	},
}
var registryFlags = map[string]string{
	"GetAstronomyPhotoOfTheDay": "date",
}
var registryShortNames = map[string]string{
	"apod": "GetAstronomyPhotoOfTheDay",
}

func Registry() map[string]FunctionWithParams {
	return registry
}

func RegistryFlags() map[string]string {
	return registryFlags
}

func RegistryShortNames() map[string]string {
	return registryShortNames
}

func GetRegistry(arg string) (FunctionWithParams, error) {
	if _, ok := registry[arg]; ok {
		return registry[arg], nil
	}
	if _, ok := registryShortNames[arg]; ok {
		return registry[registryShortNames[arg]], nil
	}
	return FunctionWithParams{}, errors.New("function not found")
}

func GetRegistryFlags(arg string) (string, error) {
	if _, ok := registryFlags[arg]; ok {
		return registryFlags[arg], nil
	}
	return "", errors.New("function not found")
}
