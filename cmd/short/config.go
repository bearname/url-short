package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type config struct {
	ServeRestAddress string
	DbAddress        string
	DbUser           string
	DbPassword       string
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

func ParseConfig() (*config, error) {
	var err error
	serveRestAddress, err := parseEnvString("SERVE_REST_ADDRESS", err)
	dbAddress, err := parseEnvString("DATABASE_ADDRESS", err)
	dbUser, err := parseEnvString("DATABASE_USER", err)
	dbPassword, err := parseEnvString("DATABASE_PASSWORD", err)

	if err != nil {
		log.Info("erro" + err.Error())
		return nil, err
	}

	return &config{
		serveRestAddress,
		dbAddress,
		dbUser,
		dbPassword,
	}, nil
}
