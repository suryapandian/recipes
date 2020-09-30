package indexer

import (
	"recipes.com/adapters"
	"recipes.com/dtos"
)

type ID struct {
	Index map[int]dtos.Recipe
}

//Creating and maintaing an integer ID for ease of operation

var idIndexer ID

func init() {
	adapters.InitializeIndex(NewIdIndexer())
}

func NewIdIndexer() *ID {
	idIndexer.Index = make(map[int]dtos.Recipe)
	return &idIndexer
}

func GetIdIndexer() *ID {
	return &idIndexer
}

func (id *ID) Populate(r dtos.Recipe, i int) {
	id.Index[i] = r
}

func (id *ID) GetResult(i *adapters.CookbookInput, r *dtos.CookbookDetails) {
	return
}
