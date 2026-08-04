package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type config struct {
	Port              string
	APIKey            string
	AllowedOrigins    []string
	MaxRequestBytes   int64
	RateLimitRPS      float64
	RateLimitBurst    int
	RequestTimeout    time.Duration
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ShutdownTimeout   time.Duration
}

// loadConfig reads configuration from the environment and fails closed:
// there is no baked-in default API key, so the service refuses to start
// rather than run with a guessable credential.
func loadConfig() (config, error) {
	cfg := config{
		Port:              getEnv("PORT", "8080"),
		APIKey:            os.Getenv("API_KEY"),
		MaxRequestBytes:   getEnvInt64("MAX_REQUEST_BYTES", 1<<20), // 1 MiB
		RateLimitRPS:      getEnvFloat("RATE_LIMIT_RPS", 5),
		RateLimitBurst:    int(getEnvInt64("RATE_LIMIT_BURST", 10)),
		RequestTimeout:    getEnvDuration("REQUEST_TIMEOUT", 10*time.Second),
		ReadTimeout:       getEnvDuration("READ_TIMEOUT", 5*time.Second),
		ReadHeaderTimeout: getEnvDuration("READ_HEADER_TIMEOUT", 5*time.Second),
		WriteTimeout:      getEnvDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:       getEnvDuration("IDLE_TIMEOUT", 60*time.Second),
		ShutdownTimeout:   getEnvDuration("SHUTDOWN_TIMEOUT", 15*time.Second),
	}

	if origins := os.Getenv("ALLOWED_ORIGINS"); origins != "" {
		for _, o := range strings.Split(origins, ",") {
			if o = strings.TrimSpace(o); o != "" {
				cfg.AllowedOrigins = append(cfg.AllowedOrigins, o)
			}
		}
	}

	if cfg.APIKey == "" {
		return cfg, fmt.Errorf("API_KEY environment variable must be set")
	}
	if len(cfg.APIKey) < 16 {
		return cfg, fmt.Errorf("API_KEY must be at least 16 characters")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			return n
		}
	}
	return fallback
}

func getEnvFloat(key string, fallback float64) float64 {
	if v := os.Getenv(key); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
