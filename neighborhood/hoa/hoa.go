// this will eventually change into a neighborhood by neighborhood config
// also, it should be integrated into deployment pipelines rather than a static config
// for now, we'll just exclude this from git
package main

import (
	// std libs

	// 3rd party libs
	"github.com/gin-gonic/gin"

	// personal libs

)

func main() {

	router := gin.Default()
	router.Run()
}