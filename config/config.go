package config

//Keeping all environment variables as a seperate package since they could be accessed and modified in other packages as need be. Also useful to set values in test cases.

import "os"

var (
	POSTCODE   string
	TIME_RANGE string
	SEARCH_STR string
)

func init() {
	POSTCODE = getEnv("POSTCODE", "10120")
	TIME_RANGE = getEnv("TIME_RANGE", "7AM-12AM")
	SEARCH_STR = getEnv("SEARCH_STR", "Chicken")
}

func getEnv(key, fallbackValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallbackValue
}
