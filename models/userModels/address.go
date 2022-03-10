package usermodels

type UserAddress struct {
	Prefecture string `json:"prefecture"`
	City       string `json:"city"`
	District   string `json:"district"`
	Street     string `json:"street"`
}
