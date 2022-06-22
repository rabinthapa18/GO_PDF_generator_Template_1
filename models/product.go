package models

type ProductData struct {
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}
