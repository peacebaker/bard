package external

import (
	"github.com/microcosm-cc/bluemonday"

	atlas "github.com/peacebaker/secretSocial/network/atlas/external"

	"github.com/peacebaker/secretSocial/errs/erresp"
)

type Session struct {
	Neighborhood string `json:"neighborhood"`
	Username     string `json:"username"`
	Token        string `json:"token"`
}

type SessionOrErr struct {
	Session Session    `json:"session"`
	Err     erresp.Msg `json:"err"`
}

func (session *Session) Sanitize() {
	policy := bluemonday.StrictPolicy()
	session.Neighborhood = policy.Sanitize(session.Neighborhood)
	session.Username = policy.Sanitize(session.Username)
	session.Token = policy.Sanitize(session.Token)
}

type Logout struct {
	Success bool `json:"success"`
}

type LogoutOrErr struct {
	Logout Logout     `json:"logout"`
	Err    erresp.Msg `json:"err"`
}

func (logout *Logout) Sanitize() {
	// bools don't need to be sanitized; they don't encode if they're invalid
}

// someday, we'll load this from a database or maybe a config file
// for now, everyone inside the neighborhood can just import this;
// backends will need to request phonebooks
// atlas will need to be manually updated if the Guard changes location
var Alpha = atlas.Phonebook{
	Chat: atlas.Server{
		Host: "localhost",
		Port: "42100",
	},
	Feed: atlas.Server{
		Host: "localhost",
		Port: "42200",
	},
	Guard: atlas.Server{
		Host: "localhost",
		Port: "42300",
	},
	HOA: atlas.Server{
		Host: "localhost",
		Port: "42400",
	},
	House: atlas.Server{
		Host: "localhost",
		Port: "42500",
	},
	Postal: atlas.Server{
		Host: "localhost",
		Port: "42600",
	},
	Sus: atlas.Server{
		Host: "localhost",
		Port: "42700",
	},
}
