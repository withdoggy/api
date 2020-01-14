package http

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func Run(dbCl *firestore.Client) {
	r := gin.Default()
	r.GET("/", listPetsHandler(dbCl))
	r.Run()

}
