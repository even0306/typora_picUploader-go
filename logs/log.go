package logs

import (
	"log"
	"os"
)

func LogFile() log.Logger {
	file := "./server.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, "[uploader]", log.Ldate|log.Ltime|log.Lshortfile)
	return *logger
}
