package Logger

import (
	"os"
	"time"
)

func WriteToLog(logFile string, data string) {
	logger, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	now := time.Now()
	if _, err := logger.WriteString(now.Format("2006-01-02 15:04:05") + " " + data + "\r\n"); err != nil {
		panic(err)
	}
}
