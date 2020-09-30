package indexer

import (
	"recipes.com/adapters"
	"recipes.com/dtos"
)

type Postcode struct {
	Index map[string][]int
}

//The key of Postcode index is the postcode of the recipe and the value is the array of recipe IDs of recipes that is deliverable to the postcode.

var postcodeIndexer Postcode

func init() {
	adapters.InitializeIndex(NewPostCodeIndex())
}

func NewPostCodeIndex() *Postcode {
	postcodeIndexer.Index = make(map[string][]int)
	return &postcodeIndexer
}

func (p *Postcode) Populate(r dtos.Recipe, i int) {
	p.Index[r.Postcode] = append(p.Index[r.Postcode], i)
}

func (p *Postcode) getPostcodeByMaxDelivery() (result dtos.BusiestPostCode) {
	result.Postcode, result.DeliveryCount = getMaxAttribute(p.Index)
	return
}

func (p *Postcode) GetResult(i *adapters.CookbookInput, r *dtos.CookbookDetails) {
	r.BusiestPostCode = p.getPostcodeByMaxDelivery()
	return
}

func GetPostcodeIndexer() *Postcode {
	return &postcodeIndexer
}

func getMaxAttribute(indexes map[string][]int) (result string, max int) {
	for r, v := range indexes {
		numberOfValues := len(v)
		if numberOfValues > max {
			max = numberOfValues
			result = r
		}
	}
	return
}
