package models

type TacoModel struct {
	IdTaco     int64   `json:"id_taco"`
	NameTaco   string  `json:"name_taco"`
	WeightTaco float64 `json:"weight_taco"`
}
