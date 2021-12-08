package graph

import "go-graphql-mdb/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	books []*model.Book
	
}
