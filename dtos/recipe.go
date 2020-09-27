package dtos

type Recipe struct {
	ID       int
	Postcode string `json:"postcode"`
	Name     string `json:"recipe"`
	Delivery string `json:"delivery"`
}
