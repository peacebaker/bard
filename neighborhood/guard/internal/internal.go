package internal

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	backend "github.com/peacebaker/secretSocial/backend/external"
	sus "github.com/peacebaker/secretSocial/neighborhood/sus/external"

	"github.com/peacebaker/secretSocial/errs"
	"github.com/peacebaker/secretSocial/errs/erresp"
	"github.com/peacebaker/secretSocial/neighborhood/guard/db"
	"github.com/peacebaker/secretSocial/neighborhood/guard/external"
)

const (
	// someday, this will vary depending upon the instance; not sure how I'll implement that yet
	NEIGHBORHOODnAME = "Alpha"

	EXPIREaFTER = 7
)

// receives a sanitized SignUp forom (originally from backend)
// verifies the request is valid
// forwards it to sus
// checks sus' response and returns it to the backend
func SignUp(form backend.SignUpForm) error {

	// try to marshall the struct back into json
	signup, err := json.Marshal(form)
	if err != nil {
		return errs.JSONEncodeFailed{Message: "failed to encode signup form into json", Err: err}
	}

	// this will probably be replaced by a db reference someday
	url := "http://" + external.Alpha.Sus.Host + ":" + external.Alpha.Sus.Port + "/signup"

	// forward the form data to sus
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(signup))
	if err != nil {
		return errs.PostReqFailed{Message: "post request to sus failed\n", Err: err}
	}
	defer resp.Body.Close()

	// wrangle sus' response back into a struct
	var response sus.SuccessOrErr
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return errs.JSONDecodeFailed{Message: "failed to decode response from sus\n", Err: err}
	}

	// return error on failure; nil on success
	if len(response.Err.Code) != 0 {
		return erresp.Code2Err(response.Err.Code)
	} else {
		return nil
	}
}

// receives a sanitized LoginForm (originally from the backend)
// verifies the request is valid
// forwards it to sus
// checks sus' response and returns it to the backend
func Login(form backend.LoginForm) error {

	// marshall the struct back into json
	login, err := json.Marshal(form)
	if err != nil {
		return errs.JSONEncodeFailed{Message: "failed to marshall login form into json\n", Err: err}
	}

	// this will probably be replaced by a db reference someday
	url := "http://" + external.Alpha.Sus.Host + ":" + external.Alpha.Sus.Port + "/signin"

	// send a request to the sus server asking for validation
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(login))
	if err != nil {
		return errs.PostReqFailed{Message: "post request to sus failed\n", Err: err}
	}
	defer resp.Body.Close()

	// wrangle sus' response back into a struct
	var response sus.Success
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return errs.JSONDecodeFailed{Message: "faled to decode response from sus\n", Err: err}
	}

	// return nil on success; error on failure
	if response.Outcome {
		return nil
	} else {

		// sus doesn't actually return a false outcome, so this code should be unreachable
		return errs.LoginFailed{Message: "login for " + form.Username + "failed\n"}
	}
}

func Logout(req external.Session) error {

	// translate the external Session into a db Session
	s := db.Session{
		Username: req.Username,
		Token:    req.Token,
	}

	// delete the session or return an error
	if err := s.Delete(); err != nil {
		return err
	}

	// silence is golden
	return nil
}

// should ONLY be run after a successful Login
// generates a new token for the given username
// saves it to the db
// returns it
func NewSession(username string) (external.Session, error) {

	// generate login token
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return external.Session{}, errs.RandSeedFailed{Message: "generation of login token seed failed\n", Err: err}
	}
	token := base64.StdEncoding.EncodeToString(bytes)

	// find expiration date
	now := time.Now()
	expires := now.AddDate(0, 0, EXPIREaFTER)

	// save the session to the database
	s := db.Session{
		Username: username,
		Token:    token,
		Expires:  expires,
	}
	if err := s.Save(); err != nil {
		return external.Session{}, err
	}

	// return the session if all goes well
	return external.Session{
		Neighborhood: NEIGHBORHOODnAME,
		Username:     s.Username,
		Token:        s.Token,
	}, nil
}

// checks if the token exists for the given user in the db
// returns true if it does, false if it doesn't
func Validate(session external.Session) bool {

	// just need to create a db session and call to validate it
	s := db.Session{
		Username: session.Username,
		Token:    session.Token,
	}
	return s.Validate()
}
