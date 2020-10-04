package adapters

import (
	"recipes.com/config"
	"recipes.com/dtos"
	"strings"
)

type ICookbook interface {
	FetchResult() dtos.CookbookDetails
}

// If we need to persist the data in a database, we could write a cookbookdb interface and simply replace the call in main.go

type Indexer interface {
	Populate(dtos.Recipe, int)
	GetResult(*CookbookInput, *dtos.CookbookDetails)
}

var Indexes = []Indexer{}

func InitializeIndex(indexer Indexer) {
	Indexes = append(Indexes, indexer)
}

/*List of indexes are initialized by init() in indexer package.
Have tried to roughly mimic RDBMS index and have created an index for each attribute like postcode, name etc....
An index here is a map with key as the postcode, name etc... and the value as a dummy recipe ID.
By creating an index like this operations are greatly minized.
Instead of looping thorugh the entire array we could simply fetch from the map
*/

type CookbookInput struct {
	Postcode      string
	TimeRange     string
	SearchQueries []string
}

func NewCookbook(recipes []dtos.Recipe) *CookbookInput {
	for i, r := range recipes {
		for _, index := range Indexes {
			index.Populate(r, i)
		}
	}
	input := CookbookInput{}
	input.Postcode = config.POSTCODE
	input.TimeRange = config.TIME_RANGE
	input.SearchQueries = strings.Split(config.SEARCH_STR, ",")
	return &input
}

func (c *CookbookInput) FetchResult() (r dtos.CookbookDetails) {
	for _, index := range Indexes {
		index.GetResult(c, &r)
	}
	return
}

/*Each index computes the part of the result that could be computed from that Index.
Fox ex: Name index would compute the number of unique recipes, since the number of unique recipes is dependent on name.
Likewise, Count of recipes for a postcode would be computed by postcode index.
*/
