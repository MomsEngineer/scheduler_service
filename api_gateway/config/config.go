package config

import (
	"log"
	"os"
)

type Config struct {
	AppointmentServiceAddr string
}

var AppConfig *Config

func Init() {
	AppConfig = &Config{
		AppointmentServiceAddr: getEnv("APPOINTMENT_SERVICE_ADDR", "localhost:50051"),
	}
	log.Println("Config initialized:", *AppConfig)
}

func getEnv(key string, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultValue
}
