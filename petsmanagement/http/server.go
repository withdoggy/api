package http

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func Run(dbCl *firestore.Client) {
	r := gin.Default()
	r.GET("/ping", listPetsHandler(dbCl))
	r.Run()

}
func AbortMsg(code int, err error, c *gin.Context) {
	c.String(code, "Oops! Please retry.")
	c.Error(err)
	c.Abort()
}
