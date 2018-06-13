package utils

import (
	// "os"
	// "io"
	// "fmt"
	log "github.com/sirupsen/logrus"

	models "CleanArchMeetingRoom/models"
)

func InitLogger() {
	 // open a file
	// f, err := os.OpenFile("logfile.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
	// }
	// // Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	// // io.MultiWriter(logf, os.Stdout).
	// log.SetOutput(io.MultiWriter(f, os.Stdout))

	// // Only log the warning severity or above.
	// log.SetLevel(log.DebugLevel)
	// return 
}

func Logger(m *models.Log) {
	contextLogger := log.WithFields(log.Fields{
		"FileName": m.FileName,
		"LineNumber": m.LineNumber,
	})
	switch m.Status {
	case "info":
		contextLogger.Info("m.Msg")
	case "error":
		contextLogger.Error("m.Msg")
	case "warn":
		contextLogger.Warn("m.Msg")
	case "fatal":
		contextLogger.Fatal("m.Msg")
	default:
		contextLogger.Error("m.Msg")
	}
}
