package main

import (
	//	"fmt"
	//	"github.com/bamzi/jobrunner"
	"log"
	"os"
)

func main() {

	myLogger("WARNING")
	//	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	//	jobrunner.Schedule("@every 5s", ReminderEmails{})
}

// Job Specific Functions
type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	//logger := myLogger()

	myLogger("INFO")
	//logger.Println("test")
	//log.Printf("Every 5 sec send reminder emails \n")
}

func myLogger(LOGLEVEL string) *log.Logger {

	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "logger:| "+LOGLEVEL, log.LstdFlags)

	logger.Println("text to append")
	logger.Println("more text to append")

	return logger

	//return logger, nil

}
