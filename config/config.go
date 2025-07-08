package config

import "os"

type Config struct {
	Port string
	MongoURI string
	MongoDB string
	Clouldinary string
}

func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "8004"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27011"),
		MongoDB: getEnv("MONGO_DB", "flash-cards"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

