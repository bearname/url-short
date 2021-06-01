package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

	"github.com/bearname/url-short/pkg/shortener/infrastructure"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("shortener.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Error(err)
			}
		}(file)
	}

	config, err := ParseConfig()
	if err != nil {
		log.Info("Default settings" + err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	 mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+config.DbAddress))

	//if done {
	//	return
	//}
	//
	//server := infrastructure.Server{}
	//killSignalChan := server.GetKillSignalChan()
	//
	//serverUrl := ":8000"
	//log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	//
	//srv := server.StartServer(serverUrl, connector)
	//
	//server.WaitForKillSignal(killSignalChan)
	//err = srv.Shutdown(context.Background())
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
}

func getConnector(err error, config *config) (config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+config.DbAddress))
}
