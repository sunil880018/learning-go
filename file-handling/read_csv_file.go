package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Product struct {
	ID         int
	Name       string
	Image      string
	Price      float64
	Quantity   int
	OutOfStock bool
}

func main() {
	file, err := os.Open("product-lambda/sample_products.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header line
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	var products []Product

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		id, _ := strconv.Atoi(record[0])
		price, _ := strconv.ParseFloat(record[3], 64)
		qty, _ := strconv.Atoi(record[4])
		outOfStock, _ := strconv.ParseBool(record[5])

		product := Product{
			ID:         id,
			Name:       record[1],
			Image:      record[2],
			Price:      price,
			Quantity:   qty,
			OutOfStock: outOfStock,
		}

		products = append(products, product)
	}

	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}
}
