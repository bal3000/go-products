package models

type ProductCollection struct {
	Products []Product `json:"products"`
}

type Product struct {
	ProductID string `json:"productId"`
	PackSizes []int  `json:"packsizes"`
}
