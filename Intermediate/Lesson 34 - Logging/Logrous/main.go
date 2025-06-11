package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	// Set a log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is a info message")
	log.Warn("This is an warn message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in.")
}
