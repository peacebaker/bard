// the guard validates backend servers and protects sus
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	backend "github.com/peacebaker/secretSocial/backend/external"

	"github.com/peacebaker/secretSocial/errs/erresp"
	"github.com/peacebaker/secretSocial/neighborhood/guard/external"
	"github.com/peacebaker/secretSocial/neighborhood/guard/internal"
	"github.com/peacebaker/secretSocial/neighborhood/guard/logger"
	"github.com/peacebaker/secretSocial/neighborhood/sus/db"
)

func SignUp(c *gin.Context) {

	// wrangle the signup form data into the signup struct
	var form backend.SignUpForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// re-sanitize, in case the backend is compromised
	form.Sanitize()

	// attempt to sign up or return error on failure
	if err := internal.SignUp(form); err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// assuming signup is successful, generate a session (save to the db) and return it
	session, err := internal.NewSession(form.Username)
	if err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}
	c.JSON(200, session)
}

func Login(c *gin.Context) {

	// wrangle login form to json and ensure the request is valid
	var form backend.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// re-sanitize in case the backend's compromised
	form.Sanitize()

	// attempt to login and return an error on failure
	if err := internal.Login(form); err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// assuming the login is successful, generate a session (saves to the db) and return it
	session, err := internal.NewSession(form.Username)
	if err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}
	c.JSON(200, session)
}

func Phonebook(c *gin.Context) {

	// wrangle the session json into a struct
	var session external.Session
	if err := c.BindJSON(session); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// re-sanitize, in case the backend gets compromised or something
	session.Sanitize()

	// return the phonebook if the token's valid
	if internal.Validate(session) {
		c.JSON(200, external.Alpha)
		return
	} else {
		logger.Log.Println("login failed for " + session.Username)
		msg := erresp.Code2Msg["Unauthorized"]
		c.JSON(403, msg)
	}
}

func Logout(c *gin.Context) {

	// wrangle request into a struct
	var session external.Session
	if err := c.BindJSON(session); err != nil {
		c.JSON(400, erresp.UWutM8)
		return
	}

	// re-sanitize, in case the backend gets compromised or something
	session.Sanitize()

	// try to logout
	if err := internal.Logout(session); err != nil {
		logger.Log.Println(err)
		status, msg := erresp.Err2Resp(err)
		c.JSON(status, msg)
		return
	}

	// silence is deadly
	success := external.Logout{Success: true}
	c.JSON(200, success)
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

	// register the routes
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/login", Login)
	router.POST("/logout", Logout)
	router.POST("/phonebook", Phonebook)
	router.POST("/signup", SignUp)

	//
	port := ":" + external.Alpha.Guard.Port
	router.Run(port)
}
