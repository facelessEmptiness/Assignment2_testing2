package config // ← Рекомендую config вместо logs (но logs тоже ОК)

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,                  // ← ОБЯЗАТЕЛЬНО!
		TimestampFormat: "2006-01-02 15:04:05", // ← ФОРМАТ, не дата!
	})

	log.SetLevel(log.InfoLevel)

	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}

	file, err := os.OpenFile("logs/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	log.Info("Logger initialized successfully")
}
