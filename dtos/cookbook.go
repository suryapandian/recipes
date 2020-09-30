package dtos

//Contracts and struct definitions that are used across packages are defined in Data Transfer Object (dto) packages.

type CookbookDetails struct {
	UniqueRecipeCount       int                     `json:"unique_recipe_count"`
	CountPerRecipe          []RecipeCount           `json:"count_per_recipe"`
	BusiestPostCode         BusiestPostCode         `json:"busiest_postcode"`
	CountPerPostCodeAndTime CountPerPostCodeAndTime `json:"count_per_postcode_and_time"`
	MatchByName             []string                `json:"match_by_name"`
}

type RecipeCount struct {
	Recipe string `json:"recipe"`
	Count  int    `json:"count"`
}

type BusiestPostCode struct {
	Postcode      string `json:"postcode"`
	DeliveryCount int    `json:"delivery_count"`
}

type CountPerPostCodeAndTime struct {
	Postcode      string `json:"postcode"`
	From          string `json:"from"`
	To            string `json:"to"`
	DeliveryCount int    `json:"delivery_count"`
}
