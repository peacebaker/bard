package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/peacebaker/secretSocial/errs"
)

// I call it SPONGEbOB case.
const (
	CLIENTuRi      = "mongodb://localhost:27017"
	DBnAME         = "vsus"
	USERcOLLECTION = "user"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var col *mongo.Collection

type User struct {
	User string `bson:"user"`
	Hash []byte `bson:"hash"`
}

// Establishes a persistent connection to the MongoDB collection.
func Connect() error {

	// make a new client w/ default options, using our uri
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(CLIENTuRi))
	if err != nil {
		return errs.MongoFailed{Message: "couldn't create mongo client", Err: err}
	}

	// still not sure what a context is
	// something about signal propagation?
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	// connect or die trying
	// service is useless without its db
	err = client.Connect(ctx)
	if err != nil {
		return errs.MongoFailed{Message: "couldn't connect to db backend", Err: err}
	}

	// collections
	col = client.Database(DBnAME).Collection(USERcOLLECTION)

	return nil
}

func Disconnect() {
	client.Disconnect(ctx)
	cancel()
}

func (u User) Save() error {

	// fail if the user already exists
	if u.Exists() {
		return errs.UserExists{Message: "the user " + u.User + "already exists in the database."}
	}

	// attempt to save
	_, err := col.InsertOne(context.Background(), u)
	if err != nil {
		return errs.MongoFailed{Message: "failed to insert the user into mongodb.", Err: err}
	}

	return nil
}

func (u User) Validate() bool {

	// create our mongo filter
	filter := bson.M{"user": u.User, "hash": u.Hash}

	// if the username and password match, return true
	var user User
	err := col.FindOne(ctx, filter).Decode(&user)

	// return false on err, true otherwise
	return err == nil
}

func (u User) Exists() bool {

	// search for the user in the db
	var user User
	filter := bson.M{"user": u.User}
	err := col.FindOne(ctx, filter).Decode(&user)

	// return false on err, true otherwise
	return err == nil
}
