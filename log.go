package main

import (
	"time"

	log "github.com/elarasu/basis/log"
	"go.uber.org/zap"
)

func init() {
	log.Configure(log.Config{
		EncodeLogsAsJson:   false,
		FileLoggingEnabled: false,
		Directory:          "/var/log/test",
		Filename:           "test.log",
		MaxSize:            20, //MB
		MaxBackups:         1,
		MaxAge:             31, //Days
	})
}

func main() {
	log.Debug("Sample debug statement")
	// use this format for non-performant logs
	logger := log.Sugar()
	// sugar api for just k,v
	logger.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	// printf style
	logger.Infof("failed to fetch URL %s", "http://sample.com")
	// use the typed fields for performant needs
	log.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	// for just string use direct log
	log.Warn("Warn a situation")
	log.Fatal("Fatal, can't proceed further")
	// this statement wouldn't be reached
	log.Error("Error")
}
