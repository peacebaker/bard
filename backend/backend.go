// The backend's primary purpose is to validate that users are logged in and route their requests inward.
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	guard "github.com/peacebaker/secretSocial/neighborhood/guard/external"
	atlas "github.com/peacebaker/secretSocial/network/atlas/external"

	"github.com/peacebaker/secretSocial/backend/external"
	"github.com/peacebaker/secretSocial/backend/internal"
	"github.com/peacebaker/secretSocial/backend/logger"
	"github.com/peacebaker/secretSocial/errs/erresp"
)

func SignUp(c *gin.Context) {

	// wrangle the signup form data into the signup struct
	var form external.SignUpForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// never rawdog the internet, AS THEY SAY!
	form.Sanitize()
	if msg := form.Validate(); len(msg.Code) != 0 {
		logger.Log.Print("Invalid Request: ", msg, "\n")
		c.JSON(200, msg)
		return
	}

	//
	session, err := internal.SignUp(form)
	if err != nil {
		logger.Log.Printf("Error type: %T\n", err)
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		logger.Log.Println(status, msg)
		c.JSON(status, msg)
		return
	}

	//
	c.JSON(200, session)
}

func Login(c *gin.Context) {

	// wrangle the login form data into the login struct
	var form external.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// AS THEY SAY!
	form.Sanitize()

	//
	session, err := internal.Login(form)
	if err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// if all goes well, pass on the session
	c.JSON(200, session)
}

func Logout(c *gin.Context) {

	//
	var session guard.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// as they say
	session.Sanitize()

	//
	logout, err := internal.Logout(session)
	if err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// if all goes well, pass on the success message
	c.JSON(200, logout)
}

// cors middleware
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		// set cors headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")

		// respond yes to preflight check
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	// open log or die trying
	if err := logger.Open(); err != nil {
		log.Fatal(err)
	}
	defer logger.Close()

	// load the GuardHouses or die trying
	if err := internal.LoadGuardHouses(); err != nil {
		log.Fatal(err)
	}

	// // I don't think we'll actually need a db for this service?
	// if err := db.Connect(); err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Disconnect()

	// fun fact: gin is peacebaker's liquor of choice
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors())

	// welcome to the neighborhood
	router.POST("/login", Login)
	router.POST("/logout", Logout)
	router.POST("/signup", SignUp)

	// start the process
	port := ":" + atlas.Backend.Port
	router.Run(port)
}
