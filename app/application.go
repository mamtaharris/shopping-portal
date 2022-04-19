package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApplication() {
	router = gin.Default()

	mappings()
	router.Run()
}
