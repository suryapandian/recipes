package indexer

import (
	"errors"
	"recipes.com/adapters"
	"recipes.com/dtos"
	"recipes.com/utils"
	"regexp"
	"strconv"
	"strings"
)

type EndTime struct {
	Index map[string][]int
}

//The key of EndTime index is the endtime in Delivery and the value is the array of recipeID that has the key as the endTime.

const TimeEndAt = 12
const TimeBeginAt = 1

// Time Range is between 1-12
//ToDo: Try and implement using go's standard time

func init() {
	adapters.InitializeIndex(NewEndTimeIndex())
}

func NewEndTimeIndex() *EndTime {
	e := EndTime{}
	e.Index = make(map[string][]int)
	return &e
}

func (e *EndTime) Populate(r dtos.Recipe, i int) {
	endTime, _ := getStartEndTimeFromDelivery(r.Delivery)
	//ToDo: Populate both StartTimeIndex and and EndTime index to avoid getStartEndTimeFromDelivery twice
	e.Index[endTime] = append(e.Index[endTime], i)
}

func (e *EndTime) GetResult(i *adapters.CookbookInput, r *dtos.CookbookDetails) {
	r.CountPerPostCodeAndTime, _ = e.GetRecipesByPostcodeTime(i.Postcode, i.TimeRange)
	return
}

func (e *EndTime) GetRecipesByPostcodeTime(postcode string, time string) (r dtos.CountPerPostCodeAndTime, err error) {

	timeString := strings.Split(time, "-")
	r.From, r.To, r.Postcode = timeString[0], timeString[1], postcode
	timeRange, err := getTimeRange(r.From, r.To)
	startTimeIndex := GetStartTimeIndexer()
	postcodeIndex := GetPostcodeIndexer()

	var startRecipes, endRecipes, result []int
	for _, time := range timeRange {
		startRecipes = append(startRecipes, startTimeIndex.Index[time]...)
		endRecipes = append(endRecipes, e.Index[time]...)
	}
	result = utils.Intersection(startRecipes, endRecipes)
	result = utils.Intersection(postcodeIndex.Index[postcode], result)
	r.DeliveryCount = len(result)
	return
}

func getTimeRange(start, end string) (r []string, err error) {
	//This functions give the list of time ranges that falls within the given start time and end time.
	re := regexp.MustCompile("[0-9]+")
	startTimeStr, endTimeStr := re.FindAllString(start, -1)[0], re.FindAllString(end, -1)[0]
	startTime, err := strconv.Atoi(startTimeStr)
	endTime, err := strconv.Atoi(endTimeStr)
	//Extracting the integer time value from the given time string

	startRange := strings.Trim(start, startTimeStr)
	endRange := strings.Trim(end, endTimeStr)
	// Extracting AM or PM from the time string

	if startRange == endRange {
		if startTime > endTime {
			err = errors.New("Invalid time range")
			return
		}
		for i := startTime; i <= endTime; i++ {
			r = append(r, strconv.Itoa(i)+startRange)
		}
		return
	}
	for i := startTime; i < TimeEndAt; i++ {
		r = append(r, strconv.Itoa(i)+startRange)
	}
	r = append(r, strconv.Itoa(TimeEndAt)+endRange)
	for i := TimeBeginAt; i <= endTime; i++ {
		r = append(r, strconv.Itoa(i)+endRange)
	}
	return
}
