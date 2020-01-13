package db

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

func NewFirestoreClient(projectId string) (*firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
func ListPets(ctx context.Context, cl *firestore.Client) (*ListPetResult, error) {
	iter := cl.Collection("pets").Documents(ctx)
	defer iter.Stop()
	var result ListPetResult
	var i int = 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var foo Pet
		doc.DataTo(&foo)
		result.Pets = append(result.Pets, foo)
		i++
	}
	result.Count = i
	return &result, nil
}

type ListPetResult struct {
	Count int   `json:"count,omitempty"`
	Pets  []Pet `json:"pets,omitempty"`
}
type Pet struct {
	Bread string `firestore:"bread,omitempty" json:"bread,omitempty"`
	Name  string `firestore:"name,omitempty" json:"name,omitempty"`
}
