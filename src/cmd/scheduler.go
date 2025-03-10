package cmd

import (
	"aion/datasources/nasa"
	"aion/pkg/logging"
	"time"

	"github.com/go-co-op/gocron"
)

func ScheduleJobs() {
	t := time.Now()
	logging.AuditLogger.Info().Msgf("Starting cron server at %s", t.Format(time.RFC3339))
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("10:00").Do(func() {
		t1 := time.Now()
		RunJob(FunctionWithParams{
			Name: "GetAstronomyPhotoOfTheDay",
			Fn:   nasa.GetAstronomyPhotoOfTheDay,
			Params: []interface{}{
				t1.Format("2006-01-02"),
			},
		})
	})
}
