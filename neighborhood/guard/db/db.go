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
	CLIENTuRi         = "mongodb://localhost:27017"
	DbNaME            = "db"
	SESSIONcOLLECTION = "session"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var col *mongo.Collection

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
	col = client.Database(DbNaME).Collection(SESSIONcOLLECTION)

	return nil
}

func Disconnect() {
	client.Disconnect(ctx)
	cancel()
}

type Session struct {
	Username string    `bson:"username"`
	Token    string    `bson:"token"`
	Expires  time.Time `bson:"expires"`
}

func (s Session) Exists() bool {

	// search for the user's session in the db
	var session Session
	filter := bson.M{"username": s.Username}
	err := col.FindOne(ctx, filter).Decode(&session)

	// return true unless there's an error
	return err == nil
}

func (s Session) Save() error {

	// I guess multiple login sessions should be allowed, right?
	// we'll just need something to clean up the old sessions

	// attempt to save
	_, err := col.InsertOne(context.Background(), s)
	if err != nil {
		return errs.MongoFailed{Message:"failed to insert the session into mongodb.", Err: err}
	}

	return nil
}

func (s Session) Delete() error {

	// create a filter to find the correct record in the collection
	var session Session
	filter := bson.M{"username": s.Username, "token": s.Token}
	err := col.FindOne(ctx, filter).Decode(&session)
	if err != nil { 
		return errs.MongoFailed{Message: "can't find the requested login token", Err: err}
	}

	// attempt to delete the record
	_, err = col.DeleteOne(context.Background(), filter)
	if err != nil {
		return errs.MongoFailed{Message: "failed to delete the requested login token", Err: err}
	}

	return nil
}

func (s Session) Validate() bool {

	// try to find the 
	var session Session
	filter := bson.M{"username": s.Username, "token": s.Token}
	err := col.FindOne(ctx, filter).Decode(&session)

	// check that the session hasn't expired and delete the record if it has
	now := time.Now()
	if now.After(session.Expires) {
		s.Delete()
		return false
	}

	// return true unless there's an error
	return err == nil
}

type Server struct {
	Host string `bson:"host"`
	Port int    `bson:"port"`
}

type Phonebook struct {
	Chat   Server `bson:"chat"`
	Feed   Server `bson:"feed"`
	Guard  Server `bson:"guard"`
	HOA    Server `bson:"hoa"`
	House  Server `bson:"house"`
	Postal Server `bson:"postal"`
	Sus    Server `bson:"sus"`
}
