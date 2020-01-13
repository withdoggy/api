package http

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/withdoggy/api/petsmanagement/db"
)

func listPetsHandler(dbCl *firestore.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := db.ListPets(context.Background(), dbCl)
		if err != nil {
			AbortMsg(500, err, c)
		}
		s, err := json.Marshal(data)
		c.JSON(200, string(s))
	}
}
