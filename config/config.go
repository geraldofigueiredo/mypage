package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppName string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret   string
	JWTExpireIn time.Duration

	LogFormat string
	DebugMode bool
}

func LoadConfig() *Config {
	// duration, err := time.ParseDuration(os.Getenv("EXPIRES_JWT"))
	// if err != nil {
	// 	panic(err)
	// }

	debugMode, err := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	if err != nil {
		panic(err)
	}

	return &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: ":" + os.Getenv("PORT"),

		// DBHost:     os.Getenv("DB_HOST"),
		// DBPort:     os.Getenv("DB_PORT"),
		// DBUser:     os.Getenv("DB_USERNAME"),
		// DBPassword: os.Getenv("DB_PASSWORD"),
		// DBName:     os.Getenv("DB_NAME"),

		// JWTSecret:   os.Getenv("SECRET_JWT"),
		// JWTExpireIn: duration,

		LogFormat: os.Getenv("LOG_FORMAT"),
		DebugMode: debugMode,
	}
}
