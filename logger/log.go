package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"server/db"
	"server/model"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	LogMutex    sync.Mutex
)

func InitLoggers() {
	infoLog, infoFile := createLogFile("info")
	errorLog, errorFile := createLogFile("error")

	InfoLogger = log.New(infoLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	log.Println("Info log file:", infoFile)
	log.Println("Error log file:", errorFile)
}

func createLogFile(logType string) (*os.File, string) {
	logPath := "logs"
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, os.ModePerm)
	}

	currentTime := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("%s_%s.log", logType, currentTime)
	logFilePath := fmt.Sprintf("%s/%s", logPath, logFileName)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not create log file:", err)
	}

	return logFile, logFileName
}

func Savelog(c *gin.Context) {
	var log model.Log
	log.Timestamp = time.Now()
	log.Username = c.GetString("jwt_username")
	log.ClientIP = c.ClientIP()
	log.RequestURI = c.Request.RequestURI
	log.Method = c.Request.Method
	log.Path = c.Request.URL.Path
	log.StatusCode = c.Writer.Status()
	LogMutex.Lock()
	defer LogMutex.Unlock()
	if err := db.GetDBlog().Create(&log).Error; err != nil {
		ErrorLogger.Printf("Failed to log register action: %v", err)
	}
}
