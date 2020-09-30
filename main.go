package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"recipes.com/adapters"
	"recipes.com/dtos"
	_ "recipes.com/indexer"
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
	fileName := "testData.json"
	if value, ok := os.LookupEnv("FILE_NAME"); ok {
		fileName = value
	}

	inputJson, err := ioutil.ReadFile("./test/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	var recipes []dtos.Recipe
	err = json.Unmarshal(inputJson, &recipes)
	if err != nil {
		log.Fatal(err)
	}

	result := newcookbook(recipes).c.FetchResult()
	resultJson, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result", string(resultJson))
}
