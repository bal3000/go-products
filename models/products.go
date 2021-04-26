package models

type ProductCollection struct {
	Products []Product `json:"products"`
}

type Product struct {
	ProductID string `json:"productId" bson:"productId,omitempty"`
	PackSizes []int  `json:"packsizes" bson:"packsizes,omitempty"`
}
