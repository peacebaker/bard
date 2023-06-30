package main

import (
	"log"

	"github.com/gin-gonic/gin"

	backend "github.com/peacebaker/secretSocial/backend/external"
	guard "github.com/peacebaker/secretSocial/neighborhood/guard/external"

	"github.com/peacebaker/secretSocial/errs/erresp"
	"github.com/peacebaker/secretSocial/neighborhood/sus/db"
	"github.com/peacebaker/secretSocial/neighborhood/sus/external"
	"github.com/peacebaker/secretSocial/neighborhood/sus/internal"
	"github.com/peacebaker/secretSocial/neighborhood/sus/logger"
)

func SignUp(c *gin.Context) {

	// wrangle the signup form data into the signup struct
	var form backend.SignUpForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// re-sanitize, in case the guard is compromised
	form.Sanitize()

	// attempt to insert the new user
	err := internal.NewUser(form.Username, form.Password)
	var res external.Success

	// if we hit an error, log the specific and throw something generic to requester
	if err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// if all goes well
	res = external.Success{
		Outcome: true,
		Info:    form.Username + " was registered successfully.",
	}
	c.JSON(200, res)
}

func Login(c *gin.Context) {

	// wrangle the login data from the supplied JSON; then re-scrub it, cause sus is paranoid
	var form backend.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}
	form.Sanitize()

	// attempt to login; if we hit an error, log the specific and throw something generic to requester
	if err := internal.Validate(form.Username, form.Password); err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// if there's no login error, it must have been successful
	res := external.Success{
		Outcome: true,
		Info:    "login succeeded",
	}
	c.JSON(200, res)
}

func main() {

	// open logs of die trying
	if err := logger.Open(); err != nil {
		log.Fatal(err)
	}
	defer logger.Close()

	// start mongodb or die trying
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()

	// register the route(s)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/signup", SignUp)
	router.POST("/signin", Login)

	// run on the postal port
	port := ":" + guard.Alpha.Sus.Port
	router.Run(port)
}
