package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	house "github.com/peacebaker/secretSocial/neighborhood/house/external"

	"github.com/peacebaker/secretSocial/errs"
)

const (
	CLIENTuRi        = "mongodb://localhost:27017"
	DbNaME           = "db"
	PERSONcOLLECTION = "people"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var personCol *mongo.Collection

func Connect() error {

	// make a new client with default options, or die trying; service is useless w/o its db
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(CLIENTuRi))
	if err != nil {
		return errs.MongoFailed{Message: "couldn't create mongo client", Err: err}
	}

	// connect to the db or die trying
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return errs.MongoFailed{Message: "couldn't connect to db backend", Err: err}
	}

	// select our collection
	personCol = client.Database(DbNaME).Collection(PERSONcOLLECTION)

	return nil
}

func Disconnect() {
	client.Disconnect(ctx)
	cancel()
}

type Person struct {
	Username string
	HouseID  int

	Friends []Person
	Posts   []house.Post
}
