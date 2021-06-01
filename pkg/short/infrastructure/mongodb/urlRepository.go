package mongodb

import (
	"context"
	"fmt"
	"github.com/bearname/url-short/pkg/short/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UrlRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUrlRepository(client *mongo.Client, collection *mongo.Collection) *UrlRepository {
	u := new(UrlRepository)
	u.client = client
	u.collection = collection
	return u
}

func (r *UrlRepository) Create(originalUrl string, customAlias string) (int64, error) {
	now := time.Now()
	expiration := now.AddDate(1, 0, 0)
	url := domain.Url{
		OriginalUrl:    originalUrl,
		CreationDate:   now.String(),
		ExpirationDate: expiration.String(),
		CustomUrl:      customAlias,
	}

	insertResult, err := r.collection.InsertOne(context.TODO(), url)
	if err != nil {
		log.Error(err)
		fmt.Println(err)
		return 0, err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return 0, err
	}

	return oid.Timestamp().Unix(), nil
}

func (r *UrlRepository) Read(objectId string) (*domain.Url, error) {

	//s := objectId.String()
	result := r.collection.FindOne(context.Background(), bson.M{"_id": objectId})
	url := domain.Url{}
	err := result.Decode(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(result)

	return &url, nil
}
