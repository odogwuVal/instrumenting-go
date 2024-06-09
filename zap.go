// // Debug will log a Debug level event
// func (log *Logger) Debug(msg string, fields ...Field)
// // Info will log an Info level event
// func (log *Logger) Info(msg string, fields ...Field)
// // Error will log an Error level event
// func (log *Logger) Error(msg string, fields ...Field)
// // With will return a logger that will log the keys and values
// specified for future log events
// func (log *Logger) With(fields ...Field) *Logger
// // Named will return a logger with a given name
// func (log *Logger) Named(s string) *Logger
package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger = logger.Named("my-app")
	logger.Info(
		"failed to fetch URL",
		zap.String("url", "https://github.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
