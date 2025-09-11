package models

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	OutOfStock bool    `json:"out_of_stock"`
}
