package logging

import (
	"aion/config"

	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var ErrorLogger zerolog.Logger
var AuditLogger zerolog.Logger

func init() {
	godotenv.Load(".env")
	auditLogFile := &lumberjack.Logger{
		Filename:   config.BaseDir + "/logs/audit.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   false,
	}
	logFile := &lumberjack.Logger{
		Filename:   config.BaseDir + "/logs/error.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   false,
	}

	ErrorLogger = zerolog.New(logFile).With().Timestamp().Logger()
	AuditLogger = zerolog.New(auditLogFile).With().Timestamp().Logger()
}
