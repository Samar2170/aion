package main

import (
	"aion/cmd"
	"aion/datasources/nasa"
	"fmt"
	"os"
	"time"

	"github.com/akamensky/argparse"
)

func RunCLI() {
	parser := argparse.NewParser("aion", "CLI for Aion")
	apodCmd := parser.NewCommand("apod", "Get Astronomy Photo of the Day")
	apodDate := apodCmd.String("d", "date", &argparse.Options{
		Required: true,
		Help:     "date",
		Default:  time.Now().Format("2006-01-02"),
	})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if apodCmd.Happened() {
		functionWithParams := cmd.FunctionWithParams{
			Name: "GetAstronomyPhotoOfTheDay",
			Fn:   nasa.GetAstronomyPhotoOfTheDay,
			Params: []interface{}{
				*apodDate,
			},
		}
		cmd.RunJob(functionWithParams)
	} else {
		fmt.Println(parser.Usage("No command provided"))
		os.Exit(1)
	}
}

func main() {
	RunCLI()
}
