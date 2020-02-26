package config

import (
	"github.com/joho/godotenv"
	"os"
	"urlshortener/utils"
)

// Configuration set up for different environments
const DefaultEnv ="local"
const PrefixEnv = ".env."

type ConfigInterface interface {
	Start()
	Get(value string) string
	Set(key string , value string) error
}

type Config struct {
	Env string
	Prefix string
}
// Return a config interface instance
func NewConfig() ConfigInterface{
	return Config{
		Env : DefaultEnv,
		Prefix: PrefixEnv,
	}
}

// Load environments variables
// If APP_ENV doesn't have any data, the default value is taken
func (c Config) Start () {
	var env = os.Getenv("APP_ENV")

	if env != "" && utils.FileExists(c.Prefix + env) {
		_ = godotenv.Load(c.Prefix + env)
		return
	}
	_ = godotenv.Load(c.Prefix + c.Env)
}
// Get env variable by value
func (c Config) Get (value string ) string {
	return os.Getenv(value)
}
// Set a env value
func (c Config) Set (key string, value string ) error {
	if err := os.Setenv(key ,value); err != nil {
		return err
	}
	return nil
}
