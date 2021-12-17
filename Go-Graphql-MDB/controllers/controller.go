package controllers

import (
	"go-graphql-mdb/database"
	"go-graphql-mdb/graph/model"
	"log"
)

const uri="YOUR DATA BASE"

const (
	DB  = "test"
	COL = "test"
)

type ControllerB interface {
	Save(book *model.Book)
	FindAll() []*model.Book
}
func Save(document *model.Book) interface{} {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := database.Connect(uri)
	if err != nil {
		panic(err)
	}
	// Free the resource when mainn dunction is  returned
	defer database.Close(client, ctx, cancel)
	cursor, err := database.SaveOne(client, ctx, DB, COL, document)
	// handle the errors.
	if err != nil {
		panic(err)
	}
	return cursor.InsertedID
}

func FindAll() []*model.Book {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := database.Connect(uri)
	if err != nil {
		panic(err)
	}
	// Free the resource when mainn dunction is  returned
	defer database.Close(client, ctx, cancel)

	cursor, err := database.Query(client, ctx, DB, COL)
	// handle the errors.
	if err != nil {
		panic(err)
	}

	var results []*model.Book
	for cursor.Next(ctx) {
		var v *model.Book
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, v)
	}
	return results
}
