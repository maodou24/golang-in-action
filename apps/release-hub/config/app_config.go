package config

import "os"

// Config contains application runtime configuration.
type Config struct {
	HTTPAddr           string
	DataPath           string
	DatabaseName       string
	DatabaseUser       string
	DatabasePassword   string
	TokenEncryptionKey string
	SessionSecret      string
}

// Load reads config from environment with sensible defaults.
func Load() Config {
	return Config{
		HTTPAddr:           getEnv("HTTP_ADDR", ":8080"),
		DataPath:           getEnv("DATA_PATH", "./data/release-hub.db"),
		DatabaseName:       getEnv("DATABASE_NAME", "release-hub"),
		DatabaseUser:       getEnv("DATABASE_USER", "release-hub"),
		DatabasePassword:   getEnv("DATABASE_PASSWORD", "release-hub"),
		TokenEncryptionKey: getEnv("TOKEN_ENCRYPTION_KEY", ""),
		SessionSecret:      getEnv("SESSION_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
