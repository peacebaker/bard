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
	BACKeNDcOLLECTION = "backend"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var col *mongo.Collection

func Connect() error {

	// make a new client with default options
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
	col = client.Database(DbNaME).Collection(BACKeNDcOLLECTION)

	// silence is golden
	return nil
}

func Disconnect() {
	client.Disconnect(ctx)
	cancel()
}
