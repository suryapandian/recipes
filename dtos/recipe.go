package dtos

type Recipe struct {
	Postcode string `json:"postcode"`
	Name     string `json:"recipe"`
	Delivery string `json:"delivery"`
}
