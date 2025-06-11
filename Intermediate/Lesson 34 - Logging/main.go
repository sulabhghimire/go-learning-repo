package main

import (
	"io"
	"log"
	"os"
)

// Standard library provides log package for logging in application
// Standard log library doesn't have feature for log level like debug, info, error however we can create custom logging functions to
// handle different levels

func main() {

	// A general log message
	log.Println("This is a log message.")

	// Setting a prefix into our log message
	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	// Adding log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Chaning the log flags
	log.Println("This is a log message with only date, time and short file name.")

	infoLogger.Println("Info log.")
	warnLogger.Println("Warn log.")
	errorLogger.Println("Error log.")
	debugLogger.Println("Debug log.")

	// Writting the logs to file instead of standard output
	// 0666 means we are granting read and write permission to everyone
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		errorLogger.Fatalln("Error opening the log file:", err)
	}
	defer logFile.Close()

	infoFileLogger := createNewFileLogger(logFile, "INFO: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	warnLogger := createNewFileLogger(logFile, "WARN: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	errorLogger := createNewFileLogger(logFile, "ERROR: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	debugLogger := createNewFileLogger(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	infoFileLogger.Println("Writing in the file logger.")
	warnLogger.Println("Writing in the file logger.")
	errorLogger.Println("Writing in the file logger.")
	debugLogger.Println("Writing in the file logger.")

}

// Creating different custom logging levels
var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
)

func createNewFileLogger(writer io.Writer, prefix string, flag int) *log.Logger {
	return log.New(writer, prefix, flag)
}

/*
Best Practices
a. Log Levels
b. Structured Logging
c. Contextual Information
d. Log Rotation
  - Technique used to manage the size of log files by perodically rotating them out and starting new ones
  - Old log data are archived or compressed or deleted so that logs are maintained well
  - Done based on file size or days passed
c. External Services
  - Graphana
  - Prometheus
  - Loki
*/
