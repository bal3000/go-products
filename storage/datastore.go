package storage

type DataStore interface {
	PackSizesForProduct(productId string) ([]int, bool, error)
}
