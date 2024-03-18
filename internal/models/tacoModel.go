package models

type TacoModel struct {
	IdTaco           int64    `json:"id_taco"`
	NameTaco         string   `json:"name_taco"`
	WeightTaco       float64  `json:"weight_taco"`
	PfcTaco          PfcInfo  `json:"pfc_taco"`
	ProductTacoNames []string `json:"product_taco_names"`
}
