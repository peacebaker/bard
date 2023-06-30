package external

import (
	"github.com/microcosm-cc/bluemonday"

	"github.com/peacebaker/secretSocial/errs/erresp"
)

const (
	MIN_USER_LENGTH = 3
	MAX_USER_LENGTH = 25

	MIN_PASS_LENGTH = 10
	MAX_PASS_LENGTH = 100

	MAX_NAME_LENGTH = 25
)

type LoginForm struct {
	Neighborhood string `json:"neighborhood"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func (form *LoginForm) Sanitize() {
	policy := bluemonday.StrictPolicy()
	form.Username = policy.Sanitize(form.Username)
	form.Password = policy.Sanitize((form.Password))
}

type SignUpForm struct {
	Neighborhood string `json:"neighborhood"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	NickName  string `json:"nickName"`
}

func (form *SignUpForm) Sanitize() {
	policy := bluemonday.StrictPolicy()
	form.Username = policy.Sanitize(form.Username)
	form.Password = policy.Sanitize((form.Password))
	form.FirstName = policy.Sanitize((form.FirstName))
	form.LastName = policy.Sanitize(form.LastName)
	form.NickName = policy.Sanitize(form.NickName)
}

func (form *SignUpForm) Validate() erresp.Msg {

	// the frontend should already do most of these, but the backend should block all bad requests regardless.
	var msg erresp.Msg
	switch {
	case len(form.Username) < MIN_USER_LENGTH:
		msg = erresp.Msg{Code: "UserTooShort", Message: "Username is too short."}

	case len(form.Username) > MAX_USER_LENGTH:
		msg = erresp.Msg{Code: "UserTooLong", Message: "Username is too long."}
	
	case len(form.Password) < MIN_PASS_LENGTH:
		msg = erresp.Msg{Code: "PassTooShort", Message: "Password is too short."}

	case len(form.Password) > MAX_PASS_LENGTH:
		msg = erresp.Msg{Code: "PassTooLong", Message: "Password is too long."}

	case len(form.FirstName) > MAX_NAME_LENGTH:
		msg = erresp.Msg{Code: "FirstNameTooLong", Message: "First name is too long."}
		
	case len(form.LastName) > MAX_NAME_LENGTH:
		msg = erresp.Msg{Code: "LastNameTooLong", Message: "Last name is too long."}

	case len(form.NickName) > MAX_NAME_LENGTH:
		msg = erresp.Msg{Code: "NickNameTooShort", Message: "Nickname is too short."}
	}

	return msg
}
