package main

import "fmt"

// PostgresConfig defines the configuration information for a postgres database
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Dialect returns the postgres dialect
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// ConnectionInfo returns the connection info for postgres configuration.
func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

// DefaultPostgresConfig defines the default postgres configuration
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postrgres",
		Password: "$PGPW",
		Name:     "meetmeup_dev",
	}
}

// Config defines the config struct
type Config struct {
	Port    int    `json:"port"`
	Env     string `json:"env"`
	Pepper  string `json:"pepper"`
	HMACKey string `json:"hmac_key"`
}

// IsProd returns True if the environement varialbe is for p roduction, false otherwise.
func (c Config) IsProd() bool {
	return c.Env == "prod"
}

// DefaultConfig returns default development configuration for the application.
func DefaultConfig() Config {
	return Config{
		Port:    8080,
		Env:     "dev",
		Pepper:  "A23094uk#aklj",
		HMACKey: "secret-key",
	}
}
