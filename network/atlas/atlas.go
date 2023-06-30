// The atlas service should accept requests from internal IP's (backends and frontends are allowed)
// Firewall should block everyone else.
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peacebaker/secretSocial/network/atlas/external"

	guard "github.com/peacebaker/secretSocial/neighborhood/guard/external"
)

// once automated deployment is complete, atlas will receive GuardHouse addresses from Poseiden/Athena
// these will be saved to the db and hopefully held in memory all of the time as well
// for now, we can't import from Guard (because guard imports from atlas), so we'll just need to manually match this
func GuardHouses(c *gin.Context) {
	GuardHouses := map[string]external.Server{
		"Alpha": guard.Alpha.Guard,
	}
	c.JSON(200, GuardHouses)
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/guardhouses", GuardHouses)
	port := ":" + external.Atlas.Port
	router.Run(port)
}
