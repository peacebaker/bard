package internal

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/peacebaker/secretSocial/errs"
	"github.com/peacebaker/secretSocial/neighborhood/sus/db"
)

// should be fine for testing
const COST = 12

// returns nil on success
func NewUser(user, pass string) error {

	// hash the password using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), COST)
	if err != nil {
		return errs.HashFailed{Message: "password hashing failed.", Err: err}
	}

	// save the new user to the database; forward the error if there is one
	newUser := db.User{User: user, Hash: hash}
	err = newUser.Save()
	if err != nil {
		return err
	}

	return nil
}

func Validate(user, pass string) error {

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), COST)
	if err != nil {
		return errs.HashFailed{Message: "password hashing failed.", Err: err}
	}

	// log the failure and return a failure if needed; otherwise, return true
	newUser := db.User{User: user, Hash: hash}
	if !newUser.Validate() {
		return errs.LoginFailed{Message: "login failed for " + user}
	}

	return nil
}
