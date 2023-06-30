package main

import (
	"log"

	"github.com/gin-gonic/gin"

	guard "github.com/peacebaker/secretSocial/neighborhood/guard/external"

	"github.com/peacebaker/secretSocial/neighborhood/chat/db"
)

func main() {

	// start mongodb or die trying
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()

	// register the route(s)
	router := gin.Default()
	router.POST("")

	// run on the postal port
	port := ":" + guard.Alpha.Postal.Port
	router.Run(port)
}
