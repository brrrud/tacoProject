package models

type RequestForCreateTaco struct {
	NameTaco string `json:"taco_name"`
	Products []int  `json:"products"`
}
