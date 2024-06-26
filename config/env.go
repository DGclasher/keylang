package config

import (
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    PublicHost string
    Port       string

    DBUser string
    DBPass string
    DBName string
    DBPort int    // Change DBPort type to int
    DBHost string
    Secret string
    SSLMode string
}

var Envs = initConfig()

func initConfig() Config {
    godotenv.Load()
    dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432")) 
    if err != nil {
        panic("Invalid DB_PORT value")
    }
    return Config{
        PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
        Port:       getEnv("PORT", "8080"),
        DBUser:     getEnv("POSTGRES_USER", "keylang"),
        DBPass:     getEnv("POSTGRES_PASSWORD", "keylang"),
        DBHost:     getEnv("POSTGRES_HOST", "localhost"),
        DBName:     getEnv("POSTGRES_DB", "keylangdb"),
        DBPort:     dbPort, 
        Secret:     getEnv("SECRET", "averylongsecretkey"),
        SSLMode:    getEnv("SSL_MODE", "disable"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
