package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/peacebaker/secretSocial/errs"
)

// I call it SPONGEbOB case.
const (
	CLIENTuRi         = "mongodb://localhost:27017"
	DbNaME            = "db"
	CHATcOLLECTION    = "chat"
	MESSAGEcOLLECTION = "message"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var chatCol *mongo.Collection
var messageCol *mongo.Collection

func Connect() error {

	// make a new client with default options, or die trying; service is useless w/o its db
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(CLIENTuRi))
	if err != nil {
		return errs.MongoFailed{Message:"couldn't create mongo client", Err: err}
	}

	// connect to the db or die trying
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return errs.MongoFailed{Message:"couldn't connect to db backend", Err: err}
	}

	// select our collection
	chatCol = client.Database(DbNaME).Collection(CHATcOLLECTION)
	messageCol = client.Database(DbNaME).Collection(MESSAGEcOLLECTION)

	return nil
}

func Disconnect() {
	client.Disconnect(ctx)
	cancel()
}
