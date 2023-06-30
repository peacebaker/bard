package internal

import (
	"bytes"
	"encoding/json"
	"net/http"

	guard "github.com/peacebaker/secretSocial/neighborhood/guard/external"
	atlas "github.com/peacebaker/secretSocial/network/atlas/external"

	"github.com/peacebaker/secretSocial/backend/external"
	"github.com/peacebaker/secretSocial/errs"
	"github.com/peacebaker/secretSocial/errs/erresp"
)

var GuardHouses map[string]atlas.Server

func LoadGuardHouses() error {

	// Grab the address of the atlas microservice.
	server := atlas.Atlas
	url := "http://" + server.Host + ":" + server.Port + "/guardhouses"

	// Send a GET request fetching a map of the guardhouses.
	resp, err := http.Get(url)
	if err != nil {
		return errs.PostReqFailed{Message: "post request to atlas failed\n", Err: err}
	}
	defer resp.Body.Close()

	// wrangle that response into our map
	if err := json.NewDecoder(resp.Body).Decode(&GuardHouses); err != nil {
		return errs.JSONDecodeFailed{Message: "failed to decode response from atlas", Err: err}
	}

	// if all goes well
	return nil
}

func Validate(session guard.Session) (atlas.Phonebook, error) {

	// find the correct GuardHouse
	guard, ok := GuardHouses[session.Neighborhood]
	if !ok {
		return atlas.Phonebook{}, errs.NeighborhoodNotFound{
			Message: "can't find a neighborhood named " + session.Neighborhood,
		}
	}
	url := "http://" + guard.Host + ":" + guard.Port + "/phonebook"

	// encode the session into json
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return atlas.Phonebook{}, errs.JSONEncodeFailed{Message: "failed to encode login session to json", Err: err}
	}

	// ask the guard server to verify the token
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(sessionJSON))
	if err != nil {
		return atlas.Phonebook{}, errs.PostReqFailed{Message: "post request to guard failed\n", Err: err}
	}
	defer resp.Body.Close()

	// check for login failure
	if resp.StatusCode != 200 {
		return atlas.Phonebook{}, errs.LoginFailed{Message: "login failed for " + session.Username}
	}

	// wrangle the phonebook into a struct
	var phonebook atlas.Phonebook
	if err := json.NewDecoder(resp.Body).Decode(&phonebook); err != nil {
		return atlas.Phonebook{}, errs.JSONDecodeFailed{
			Message: "couldn't decode phonebook for " + session.Neighborhood,
			Err:     err,
		}
	}

	// if all goes well
	return phonebook, nil
}

func SignUp(form external.SignUpForm) (guard.Session, error) {

	// find the requested GuardHouse
	gh, ok := GuardHouses[form.Neighborhood]
	if !ok {
		return guard.Session{}, errs.NeighborhoodNotFound{
			Message: "can't find a neighborhood named " + form.Neighborhood,
		}
	}
	url := "http://" + gh.Host + ":" + gh.Port + "/signup"

	// encode the form back into json
	formJSON, err := json.Marshal(form)
	if err != nil {
		return guard.Session{}, errs.JSONEncodeFailed{Message: "failed to encode signup data to json", Err: err}
	}

	// ask the guard server to pass on the SignUp form to sus for user creation.
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(formJSON))
	if err != nil {
		return guard.Session{}, errs.PostReqFailed{Message: "internal post request failed", Err: err}
	}
	defer resp.Body.Close()

	// wrangle the response
	var response guard.SessionOrErr
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return guard.Session{}, errs.JSONDecodeFailed{Message: "failed to decode response from guard\n", Err: err}
	}

	// check for an error further upstream
	if len(response.Err.Code) != 0 {
		return guard.Session{}, erresp.Code2Err(response.Err.Code)
	} else {
		return response.Session, nil
	}
}

func Login(form external.LoginForm) (guard.Session, error) {

	// find the requested GuardHouse
	gh, ok := GuardHouses[form.Neighborhood]
	if !ok {
		return guard.Session{}, errs.NeighborhoodNotFound{Message: "can't find a neighborhood named " + form.Neighborhood}
	}
	url := "http://" + gh.Host + ":" + gh.Port + "/login"

	// encode the form back into json
	formJSON, err := json.Marshal(form)
	if err != nil {
		return guard.Session{}, errs.JSONEncodeFailed{Message: "failed to encode the login data to json", Err: err}
	}

	// ask the guard server to pass on the LogIn form to sus for user creation.
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(formJSON))
	if err != nil {
		return guard.Session{}, errs.PostReqFailed{Message: "internal post request failed", Err: err}
	}
	defer resp.Body.Close()

	// wrangle the response
	var response guard.SessionOrErr
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return guard.Session{}, errs.JSONDecodeFailed{Message: "failed to decode response from guard", Err: err}
	}

	// check for an error further upstream
	if len(response.Err.Code) != 0 {
		return guard.Session{}, erresp.Code2Err(response.Err.Code)
	} else {
		return response.Session, nil
	}
}

func Logout(session guard.Session) (guard.Logout, error) {

	// find the requested GuardHouse
	gh, ok := GuardHouses[session.Neighborhood]
	if !ok {
		return guard.Logout{}, errs.NeighborhoodNotFound{
			Message: "can't find a neighborhood named " + session.Neighborhood,
		}
	}
	url := "http://" + gh.Host + ":" + gh.Port + "/logout"

	// encode the session back to JSON
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return guard.Logout{}, errs.JSONEncodeFailed{Message: "failed to encode the session data to json", Err: err}
	}

	// ask the guard server to log out
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(sessionJSON))
	if err != nil {
		return guard.Logout{}, errs.PostReqFailed{Message: "internal post request failed", Err: err}
	}
	defer resp.Body.Close()

	// wrangle the response
	var response guard.LogoutOrErr
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return guard.Logout{}, errs.JSONDecodeFailed{Message: "failed to decode response from guard", Err: err}
	}

	// check for upstream errors or return
	if len(response.Err.Code) != 0 {
		return guard.Logout{}, erresp.Code2Err(response.Err.Code)
	} else {
		return response.Logout, nil
	}
}
