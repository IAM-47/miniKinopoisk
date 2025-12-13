package models

type Budget_and_Fees struct {
	ID                int     `json:"id"`
	IDMovie           int     `json:"id_movie"`
	TotalBudget       float64 `json:"total_budget,omitempty"`
	FeesInProdCountry float64 `json:"fees_in_prod_country,omitempty"`
	FeesInOther       float64 `json:"fees_in_other,omitempty"`
}
