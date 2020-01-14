package http

import (
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

func getDictionariesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, span := trace.StartSpan(c, "handlers.getdictionaries")
		defer span.End()
		c.JSON(200, string("s"))
	}
}

func abortMsg(code int, err error, c *gin.Context) {
	c.String(code, "Oops! Please retry.")
	c.Error(err)
	c.Abort()
}
