package http

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/", getDictionariesHandler())
	r.Run()

}
