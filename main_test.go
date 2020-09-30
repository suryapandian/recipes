package main

import (
	"recipes.com/config"
	"recipes.com/dtos"
	_ "recipes.com/indexer"
	"reflect"
	"testing"
)

func TestFetchResult(t *testing.T) {
	recipes := getDummyRecipe()

	config.SEARCH_STR = "Dosa,Ghee"
	config.POSTCODE = "10137"
	config.TIME_RANGE = "5AM-1PM"

	result := newcookbook(recipes).c.FetchResult()

	if result.UniqueRecipeCount != 4 {
		t.Errorf("Unique Recipe Count got %d, want %d", result.UniqueRecipeCount, 4)
	}
	for _, r := range result.CountPerRecipe {
		if r.Recipe == "Ghee Dosa" && r.Count != 2 {
			t.Errorf("Count per recipe for %s got %d, want %d", r.Recipe, r.Count, 2)
		}
	}
	if result.BusiestPostCode.Postcode != "10137" {
		t.Errorf("Busiest Postcode got %s, want %s", result.BusiestPostCode.Postcode, "10137")
	}
	if result.BusiestPostCode.DeliveryCount != 4 {
		t.Errorf("Busiest Postcode count got %d, want %d", result.BusiestPostCode.DeliveryCount, 4)
	}
	if result.CountPerPostCodeAndTime.DeliveryCount != 3 {
		t.Errorf("Count per postcode %s between %s got %d, want %d", config.POSTCODE, config.TIME_RANGE, result.CountPerPostCodeAndTime.DeliveryCount, 3)
	}
	if !reflect.DeepEqual(result.MatchByName, []string{"Ghee Dosa", "Masala Dosa"}) {
		t.Errorf("Recipes matched for search query %s got %d, want %d", config.SEARCH_STR, result.CountPerPostCodeAndTime.DeliveryCount, 2)
	}
}

func getDummyRecipe() []dtos.Recipe {
	r1 := dtos.Recipe{}
	r1.Postcode = "10137"
	r1.Name = "Ghee Dosa"
	r1.Delivery = "Sunday 6AM - 12PM"

	r2 := dtos.Recipe{}
	r2.Postcode = "10599"
	r2.Name = "Chapathi"
	r2.Delivery = "Sunday 11AM - 5PM"

	r3 := dtos.Recipe{}
	r3.Postcode = "10137"
	r3.Name = "Masala Dosa"
	r3.Delivery = "Sunday 7PM - 10PM"

	r4 := dtos.Recipe{}
	r4.Postcode = "10137"
	r4.Name = "Ghee Dosa"
	r4.Delivery = "Sunday 6AM - 1PM"

	r5 := dtos.Recipe{}
	r5.Postcode = "10137"
	r5.Name = "Vadai"
	r5.Delivery = "Sunday 5AM - 11AM"
	result := []dtos.Recipe{r1, r2, r3, r4, r5}
	return result
}
