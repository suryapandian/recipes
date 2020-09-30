package indexer

import (
	"recipes.com/adapters"
	"recipes.com/dtos"
	"sort"
	"strings"
)

type Item struct {
	Index map[string][]int
}

//The key of Item index is an Item from the RecipeName and the value is the array of recipeID that has the key as part of the recipe name.

func init() {
	adapters.InitializeIndex(NewItemIndex())
}

func NewItemIndex() *Item {
	n := Item{}
	n.Index = make(map[string][]int)
	return &n
}

func (i *Item) Populate(r dtos.Recipe, id int) {
	words := strings.Fields(r.Name)
	for _, word := range words {
		i.Index[word] = append(i.Index[word], id)
	}

}
func (i *Item) search(names []string) (result []string) {
	recipeList := make(map[string]bool) // To avoid duplicates in result, populating in temp map
	idIndexer := GetIdIndexer()
	for _, name := range names {
		for _, recipeID := range i.Index[name] {
			recipeName := idIndexer.Index[recipeID].Name
			if _, ok := recipeList[recipeName]; !ok {
				result = append(result, recipeName)
				recipeList[recipeName] = true
			}
		}
	}

	sort.Strings(result)
	return
}
func (i *Item) GetResult(input *adapters.CookbookInput, r *dtos.CookbookDetails) {
	r.MatchByName = i.search(input.SearchQueries)
	return
}
