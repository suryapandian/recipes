package adapters

import (
	"errors"
	"fmt"
	"recipes.com/dtos"
	"recipes.com/utils"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ICookbook interface {
	GetNumOfUniqueRecipes() (int, error)
	GetCountByRecipeName() ([]dtos.RecipeCount, error)
	GetPostcodeByMaxDelivery() (dtos.BusiestPostCode, error)
	GetRecipesByPostcodeTime(postcode string, time string) (dtos.CountPerPostCodeAndTime, error)
	SearchByName(name string) ([]string, error)
}

//var Cookbook []interface{}

type Cookbook struct {
	IndexByID        map[int]dtos.Recipe
	IndexByName      map[string][]int
	IndexByPostcode  map[string][]int
	IndexByItem      map[string][]int
	IndexByStartTime map[string][]int
	IndexByEndTime   map[string][]int
}

func NewCookbook(recipes []dtos.Recipe) *Cookbook {
	c := Cookbook{}
	c.IndexByID = make(map[int]dtos.Recipe)
	c.IndexByName = make(map[string][]int)
	c.IndexByPostcode = make(map[string][]int)
	c.IndexByItem = make(map[string][]int)
	c.IndexByStartTime = make(map[string][]int)
	c.IndexByEndTime = make(map[string][]int)

	for i, r := range recipes {
		c.IndexByID[i] = r
		c.IndexByName[r.Name] = append(c.IndexByName[r.Name], i)
		c.IndexByPostcode[r.Postcode] = append(c.IndexByPostcode[r.Postcode], i)
		words := strings.Fields(r.Name)
		for _, word := range words {
			c.IndexByItem[word] = append(c.IndexByItem[word], i)
		}
		startTime, endTime := c.getStartEndTimeFromDelivery(r.Delivery)
		c.IndexByStartTime[startTime] = append(c.IndexByStartTime[startTime], i)
		c.IndexByEndTime[endTime] = append(c.IndexByEndTime[endTime], i)
	}
	return &c
}

func (c *Cookbook) GetNumOfUniqueRecipes() (int, error) {
	return len(c.IndexByID), nil
}

func (c *Cookbook) GetCountByRecipeName() (result []dtos.RecipeCount, err error) {
	for r, v := range c.IndexByName {
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

func (c *Cookbook) GetPostcodeByMaxDelivery() (result dtos.BusiestPostCode, err error) {
	result.Postcode, result.DeliveryCount, err = c.getMaxAttribute(c.IndexByPostcode)
	return
}

func (c *Cookbook) getMaxAttribute(indexes map[string][]int) (result string, max int, err error) {
	for r, v := range indexes {
		numberOfValues := len(v)
		if numberOfValues > max {
			max = numberOfValues
			result = r
		}
	}
	return
}

func (c *Cookbook) SearchByName(names string) (result []string, err error) {
	recipeList := make(map[string]bool)
	for _, name := range strings.Split(names, ",") {
		for _, recipeID := range c.IndexByItem[name] {
			recipeName := c.IndexByID[recipeID].Name
			if _, ok := recipeList[recipeName]; !ok {
				result = append(result, recipeName)
				recipeList[recipeName] = true
			}
		}
	}
	sort.Strings(result)
	return
}

func (c *Cookbook) GetRecipesByPostcodeTime(postcode string, time string) (r dtos.CountPerPostCodeAndTime, err error) {

	timeString := strings.Split(time, "-")
	r.From, r.To, r.Postcode = timeString[0], timeString[1], postcode
	timeRange, err := c.getTimeRange(r.From, r.To)

	var startRecipes, endRecipes, result []int
	for _, time := range timeRange {
		startRecipes = append(startRecipes, c.IndexByStartTime[time]...)
		endRecipes = append(endRecipes, c.IndexByEndTime[time]...)
	}
	result = utils.Intersection(startRecipes, endRecipes)
	result = utils.Intersection(c.IndexByPostcode[postcode], result)
	r.DeliveryCount = len(result)
	return
}

func (c *Cookbook) getTimeRange(start, end string) (r []string, err error) {

	re := regexp.MustCompile("[0-9]+")
	startTimeStr, endTimeStr := re.FindAllString(start, -1)[0], re.FindAllString(end, -1)[0]
	startTime, err := strconv.Atoi(startTimeStr)
	endTime, err := strconv.Atoi(endTimeStr)

	startRange := strings.Trim(start, startTimeStr)
	endRange := strings.Trim(end, endTimeStr)

	if startRange == endRange {
		if startTime > endTime {
			fmt.Println("Invalid time range", start, end)
			err = errors.New("Invalid time range")
			return
		}
		for i := startTime; i <= endTime; i++ {
			r = append(r, string(startTime)+startRange)
		}
		return
	}
	for i := startTime; i < 12; i++ {
		r = append(r, strconv.Itoa(i)+startRange)
	}
	r = append(r, `12`+endRange)
	for i := 1; i <= endTime; i++ {
		r = append(r, strconv.Itoa(i)+endRange)
	}
	fmt.Println(r)
	return
}

func (c *Cookbook) getStartEndTimeFromDelivery(delivery string) (startTime, endTime string) {
	dayTime := strings.Fields(delivery)
	startTime = dayTime[1]
	endTime = dayTime[3]
	return
}
