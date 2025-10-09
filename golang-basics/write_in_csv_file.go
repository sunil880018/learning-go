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
	// Read from CSV
	file, err := os.Open("product-lambda/sample_products.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header
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

	// Write to new CSV file
	outFile, err := os.Create("product-lambda/output_products.csv")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)

	// Write header
	header := []string{"id", "name", "image", "price", "quantity", "out_of_stock"}
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	// Write product data
	for _, p := range products {
		record := []string{
			strconv.Itoa(p.ID),
			p.Name,
			p.Image,
			fmt.Sprintf("%.2f", p.Price),
			strconv.Itoa(p.Quantity),
			strconv.FormatBool(p.OutOfStock),
		}
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}

	fmt.Println("✅ Successfully wrote products to output_products.csv")
}
