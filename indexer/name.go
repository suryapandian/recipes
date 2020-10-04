package indexer

import (
	"recipes.com/adapters"
	"recipes.com/dtos"
	"sort"
)

type Name struct {
	Index map[string][]int
}

//The key of Name index is the name of the recipe and the value is the array of recipeID with the recipeName as that of the key.

func init() {
	adapters.InitializeIndex(NewNameIndex())
}

func NewNameIndex() *Name {
	n := Name{}
	n.Index = make(map[string][]int)
	return &n
}

func (n *Name) Populate(r dtos.Recipe, i int) {
	n.Index[r.Name] = append(n.Index[r.Name], i)
}

func (n *Name) getCountByName() (result []dtos.RecipeCount) {
	for r, v := range n.Index {
		recipe := dtos.RecipeCount{
			Recipe: r,
			Count:  len(v),
		}
		result = append(result, recipe)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Recipe < result[j].Recipe
	})
	return
}

func (n *Name) GetResult(i *adapters.CookbookInput, r *dtos.CookbookDetails) {
	r.UniqueRecipeCount = len(n.Index)
	r.CountPerRecipe = n.getCountByName()
	return
}
