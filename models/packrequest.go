package models

type PackRequest struct {
	ProductID  string `json:"productId"`
	ItemsOrder int    `json:"itemsOrder"`
}
