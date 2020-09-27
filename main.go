package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"recipes.com/adapters"
	"recipes.com/dtos"
)

type cookbook struct {
	c adapters.ICookbook
}

func newcookbook(recipes []dtos.Recipe) *cookbook {
	return &cookbook{
		c: adapters.NewCookbook(recipes),
	}
}

func main() {
	//input, err := ioutil.ReadFile("/Users/suryapandian/Downloads/hf.json")
	fileName := flag.String("file", "testData.json", "fixture file path")
	searchStrs := flag.String("search", "Steak", "search terms")
	postcode := flag.String("postcode", "10158", "postcode")
	timeRange := flag.String("time", "9AM-4PM", "time range")
	flag.Parse()
	fmt.Println(*searchStrs, *fileName, *timeRange)

	input, err := ioutil.ReadFile("./test/" + *fileName)
	if err != nil {
		fmt.Println("error while reading file")
	}

	recipes := []dtos.Recipe{}
	err = json.Unmarshal(input, &recipes)
	if err != nil {
		fmt.Println("error while reading file")
	}
	fmt.Println("len", len(recipes))

	cookbook := newcookbook(recipes)
	result := dtos.CookbookDetails{}
	result.UniqueRecipeCount, _ = cookbook.c.GetNumOfUniqueRecipes()
	result.BusiestPostCode, _ = cookbook.c.GetPostcodeByMaxDelivery()
	result.CountPerRecipe, _ = cookbook.c.GetCountByRecipeName()
	result.MatchByName, _ = cookbook.c.SearchByName(*searchStrs)
	result.CountPerPostCodeAndTime, _ = cookbook.c.GetRecipesByPostcodeTime(*postcode, *timeRange)

	d, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Errorf("Error while marshalling json %v \n", err)
	}
	fmt.Println("result", string(d))

}
