package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	// set this if environment is production
	log.SetFormatter(&log.JSONFormatter{})
	// set this for non prod
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	//fmt.Println("log")
	log.Debug("debug statement")
	log.Info("info text")
	log.Warn("Warn a situation")
	log.Error("Error")
	log.Fatal("Fatal, can't proceed further")
}
