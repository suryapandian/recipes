package indexer

import (
	"recipes.com/adapters"
	"recipes.com/dtos"
	"strings"
)

type StartTime struct {
	Index map[string][]int
}

//The key of StartTime index is the startTime in Delivery and the value is the array of recipeID that has the key as the startTime.

var startTimeIndexer StartTime

func init() {
	adapters.InitializeIndex(NewStartTimeIndex())
}

func NewStartTimeIndex() *StartTime {
	startTimeIndexer.Index = make(map[string][]int)
	return &startTimeIndexer
}

func (s *StartTime) Populate(r dtos.Recipe, i int) {
	sartTime, _ := getStartEndTimeFromDelivery(r.Delivery)
	s.Index[sartTime] = append(s.Index[sartTime], i)
}

func (s *StartTime) GetResult(i *adapters.CookbookInput, r *dtos.CookbookDetails) {
	return
}

func GetStartTimeIndexer() *StartTime {
	return &startTimeIndexer
}

func getStartEndTimeFromDelivery(delivery string) (startTime, endTime string) {
	dayTime := strings.Fields(delivery)
	startTime = dayTime[1]
	endTime = dayTime[3]
	return
}
