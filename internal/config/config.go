package config

import (
	"flag"
	"os"
	"time"
)

type Config struct {
	DatabaseURI     string
	EndPointAddress string
	TokenExp        time.Duration
	TokenSecretKey  string
	MigrationsPath  string
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.parseConfig()
	return cfg
}

func (c *Config) parseConfig() {
	c.parseFlags()
	c.parseEnv()
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.DatabaseURI, "d", "", "uri for connect to postgres database")
	flag.StringVar(&c.EndPointAddress, "a", "127.0.0.1:8080", "uri for run server")
	flag.Parse()
}

func (c *Config) parseEnv() {
	if uri := os.Getenv("TAXI_DRIVE_DATABASE_URI"); uri != "" && c.DatabaseURI == "" {
		c.DatabaseURI = uri
	}

	if runAddr := os.Getenv("TAXI_DRIVE_RUN_ADDRESS"); runAddr != "" && c.EndPointAddress == "127.0.0.1:8080" {
		c.EndPointAddress = runAddr
	}

	if tokenExp := os.Getenv("TAXI_DRIVE_TOKEN_EXP"); tokenExp != "" && c.TokenExp == time.Hour*3 {
		time, err := time.ParseDuration(tokenExp)
		if err == nil {
			c.TokenExp = time
		}
	}

	if tokenSecret := os.Getenv("TAXI_DRIVE_TOKEN_SECRET_KEY"); tokenSecret != "" {
		c.TokenSecretKey = tokenSecret
	}

	if migrationsPath := os.Getenv("TAXI_DRIVE_MIGRATIONS_PATH"); migrationsPath != "" {
		c.MigrationsPath = migrationsPath
	}
}
