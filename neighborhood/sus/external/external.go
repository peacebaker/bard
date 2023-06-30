package external

import "github.com/peacebaker/secretSocial/errs/erresp"

type Success struct {
	Outcome bool   `json:"outcome"`
	Info    string `json:"info"`
}

type SuccessOrErr struct {
	Success Success    `json:"response"`
	Err     erresp.Msg `json:"err"`
}
