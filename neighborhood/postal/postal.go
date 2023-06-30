package main

import (

	"github.com/gin-gonic/gin"
	guard "github.com/peacebaker/secretSocial/neighborhood/guard/external"

)

func main() {

	// register the route(s)
	router := gin.Default()

	// run on the postal port
	port := ":" + guard.Alpha.Postal.Port
	router.Run(port)
}