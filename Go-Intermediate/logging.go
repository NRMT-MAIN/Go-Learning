package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func loggingExample() {
	log.Println("This is a log message")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is log message with only date") 

	infoLogger.Println("This is an info message.")
	warnLogger.Println("This is a warning message.")
	errorLogger.Println("This is an error message.")

	//Adding Logs in file
	file , err := os.OpenFile("app.log" , os.O_CREATE | os.O_WRONLY | os.O_APPEND , 0666)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open log file: " , err)
	}

	debugLogger := log.New(file , "DEBUG: " , log.Ldate | log.Ltime | log.Lshortfile)

	debugLogger.Println("This is an debug message")

	// logrus library

	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is an info message.")
	log.Warn("This is an error message.")
	log.Error("This is an error message.")
	
	log.WithFields(logrus.Fields{
		"username" : "Nirmit Sahu" , 
		"method" : "GET" , 
	}).Info("User logged in.")

	logger , err := zap.NewProduction()

	if err != nil {
		log.Println("Error in initializing Zap Logger")
	}

	defer logger.Sync()

	logger.Info("This is an info message.")

	logger.Info("User looged in" , zap.String("username" , "Nirmit") , zap.String("method" , "GET"))
}

var (
	infoLogger = log.New(os.Stdout , "INFO: " , log.Ldate | log.Ltime | log.Lshortfile)
	warnLogger = log.New(os.Stdout , "WARN: " , log.Ldate | log.Ltime | log.Lshortfile)
	errorLogger = log.New(os.Stdout , "ERROR: " , log.Ldate | log.Ltime | log.Lshortfile)
)