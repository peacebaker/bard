package external

import "time"

type Person struct {
	Username	string		`json:"username"`
	Token		string		`json:"token"`
	Expires		time.Time	`json:"expires"`

	FeedID	int		`json:"feedid"`
	ChatIDs []int	`json:"chatids"`

	First	string	`json:"first"`
	Last	string	`json:"last"`
	Nick	string	`json:"nick"`
}