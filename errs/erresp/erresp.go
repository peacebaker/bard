package erresp

import (
	"github.com/peacebaker/secretSocial/errs"
)

// this is common enough to have its own variable
var UWutM8 = Msg{
	Code:    "UWutM8",
	Message: "didn't understand the request",
}

// Msg is the only error response thrown externally by secretSocial
// for the most part, every microservice will deal with or at least log its own problems,
// but if it's unable to, it may return a simple error.
// Regardless of the reason, the requesting microservice should be able to continue functioning as normal,
// even if it has to return an error to its requester.
type Msg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ruhRoh is our "uncaught error" message
type ruhRoh struct {
	message string
}

func (e ruhRoh) Error() string {
	return e.message
}

// Code2Msg
var Code2Msg = map[string]Msg{
	"UWutM8": UWutM8,

	"HashFailed": {
		Code:    "HashFailed",
		Message: "password hashing failed",
	},
	"JSONDecodeFailed": {
		Code:    "JSONDecodeFailed",
		Message: "json decoding failed",
	},
	"JSONEncodeFailed": {
		Code:    "JSONEncodeFailed",
		Message: "json encoding failed",
	},
	"LoginFailed": {
		Code:    "LoginFailed",
		Message: "login failed",
	},
	"MongoFailed": {
		Code:    "MongoFailed",
		Message: "mongodb failure",
	},
	"NeighborhoodNotFound": {
		Code:    "NeighborhoodNotFound",
		Message: "the requested neighborhood was not found",
	},
	"PostReqFailed": {
		Code:    "PostReqFailed",
		Message: "a post request between microservices has failed",
	},
	"RandSeedFailed": {
		Code:    "RandSeedFailed",
		Message: "random seed generation failed",
	},
	"Unauthorized": {
		Code:    "Unauthorized",
		Message: "authentication failed",
	},
	"UserExists": {
		Code:    "UserExists",
		Message: "a user by that name already exists",
	},

	"RuhRoh": {
		Code:    "RuhRoh",
		Message: "this shouldn't have happened",
	},
}

func Err2Resp(err error) (int, Msg) {

	// messages
	switch err.(type) {

	case errs.UWutM8:
		return 400, Code2Msg["UWutM8"]

	case errs.HashFailed:
		return 500, Code2Msg["HashFailed"]

	case errs.JSONDecodeFailed:
		return 500, Code2Msg["JSONDecodeFailed"]

	case errs.JSONEncodeFailed:
		return 500, Code2Msg["JSONEncodeFailed"]

	case errs.LoginFailed:
		return 200, Code2Msg["LoginFailed"]

	case errs.MongoFailed:
		return 500, Code2Msg["MongoFailed"]

	case errs.NeighborhoodNotFound:
		return 200, Code2Msg["NeighborhoodNotFound"]

	case errs.PostReqFailed:
		return 500, Code2Msg["PostReqFailed"]

	case errs.RandSeedFailed:
		return 500, Code2Msg["RandSeedFailed"]

	case errs.Unauthorized:
		return 403, Code2Msg["Unauthorized"]

	case errs.UserExists:
		return 200, Code2Msg["UserExists"]

	// this should... hopefully, never happen.
	default:
		return 500, Code2Msg["RuhRoh"]
	}
}

func Code2Err(code string) error {

	switch code {

	case "UWutM8":
		return errs.UWutM8{}

	case "HashFailed":
		return errs.HashFailed{}

	case "JSONDecodeFailed":
		return errs.JSONDecodeFailed{}

	case "JSONEncodeFailed":
		return errs.JSONEncodeFailed{}

	case "LoginFailed":
		return errs.LoginFailed{}

	case "MongoFailed":
		return errs.MongoFailed{}

	case "NeighborhoodNotFound":
		return errs.NeighborhoodNotFound{}

	case "PostReqFailed":
		return errs.PostReqFailed{}

	case "RandSeedFailed":
		return errs.RandSeedFailed{}

	case "Unauthorized":
		return errs.Unauthorized{}

	case "UserExists":
		return errs.UserExists{}
	}

	return ruhRoh{message: "this shouldn't have happened"}
}
