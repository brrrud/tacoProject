package models

type ProductModel struct {
	NameProduct   string  `json:"name_product"`
	WeightProduct float64 `json:"weight_product"`
	PfcProduct    PfcInfo `json:"pfc_product"`
}
