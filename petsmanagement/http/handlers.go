package http

import (
	"encoding/json"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/withdoggy/api/petsmanagement/db"
	"go.opencensus.io/trace"
)

func listPetsHandler(dbCl *firestore.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := trace.StartSpan(c, "handlers.getPets")
		defer span.End()
		data, err := db.ListPets(ctx, dbCl)
		if err != nil {
			abortMsg(500, err, c)
		}
		s, err := json.Marshal(data)
		c.JSON(200, string(s))
	}
}
func abortMsg(code int, err error, c *gin.Context) {
	c.String(code, "Oops! Please retry.")
	c.Error(err)
	c.Abort()
}
