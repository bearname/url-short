package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	ServeRestAddress string
	DbAddress        string
	DbName           string
	DbUser           string
	DbPassword       string
	MaxConnections   int
	AcquireTimeout   int
}

func parseEnvString(key string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	str, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf("undefined environment variable %v", key))
	}
	return str, nil
}

func parseEnvInt(key string, err error) (int, error) {
	s, err := parseEnvString(key, err)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func ParseConfig() (*Config, error) {
	var err error
	serveRestAddress, err := parseEnvString("SERVE_REST_ADDRESS", err)
	dbAddress, err := parseEnvString("DATABASE_ADDRESS", err)
	dbName, err := parseEnvString("DATABASE_NAME", err)
	dbUser, err := parseEnvString("DATABASE_USER", err)
	dbPassword, err := parseEnvString("DATABASE_PASSWORD", err)
	maxConnections, err := parseEnvInt("DATABASE_MAX_CONNECTION", err)
	acquireTimeout, err := parseEnvInt("DATABASE_ACQUIRE_TIMEOUT", err)

	if err != nil {
		log.Info("error " + err.Error())
		return nil, err
	}

	return &Config{
		serveRestAddress,
		dbAddress,
		dbName,
		dbUser,
		dbPassword,
		maxConnections,
		acquireTimeout,
	}, nil
}
