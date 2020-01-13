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
		ctx, span := trace.StartSpan(c, "handlers.listpets")
		defer span.End()
		data, err := db.ListPets(ctx, dbCl)
		if err != nil {
			AbortMsg(500, err, c)
		}
		s, err := json.Marshal(data)
		c.JSON(200, string(s))
	}
}
