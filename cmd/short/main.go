package main

import (
	"context"
	"fmt"
	"github.com/bearname/url-short/pkg/short/infrastructure"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	//"github.com/bearname/url-short/pkg/short/infrastructure"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	//file, err := os.OpenFile("short.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//if err == nil {
	//	log.SetOutput(file)
	//	defer func(file *os.File) {
	//		err := file.Close()
	//		if err != nil {
	//			log.Error(err)
	//		}
	//	}(file)
	//}

	//conf, err := ParseConfig()
	//if err != nil {
	//	log.Fatal("Default settings" + err.Error())
	//}
	conf := Config{
		":8080",
		"localhost:27017",
		"url-short",
		"url-short",
		"1234",
	}

	client, err := getConnector(&conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database(conf.DbName).Collection("trainers")
	//
	//ash := Trainer{"Ash", 10, "Pallet Town"}
	//insertResult, err := collection.InsertOne(context.TODO(), ash)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//var result Trainer
	//filter := bson.D{{"name", "Ash"}}
	//err = collection.FindOne(context.TODO(), filter).Decode(&result)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Printf("Found a single document: %+v\n", result)

	server := infrastructure.Server{}
	killSignalChan := server.GetKillSignalChan()

	serverUrl := ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")

	srv := server.StartServer(serverUrl, client, collection)

	server.WaitForKillSignal(killSignalChan)
	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Error(err)
		return
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err = client.Disconnect(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(client, context.TODO())
	fmt.Println("Connection to MongoDB closed.")
}

func getConnector(config *Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := "mongodb://" + config.DbUser + ":" + config.DbPassword + "@" +
		config.DbAddress + "/" + config.DbName + "?authSource=admin"
	u := "mongodb://url-short:1234@localhost:27017/url-short?authSource=admin"
	fmt.Println(uri == u)
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}
