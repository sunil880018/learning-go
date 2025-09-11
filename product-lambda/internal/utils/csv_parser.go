package utils

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"product-lambda/internal/models"
)

func ParseCSV(filePath string) ([]models.Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var products []models.Product

	_, _ = reader.Read() // skip header

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		price, _ := strconv.ParseFloat(row[3], 64)
		qty, _ := strconv.Atoi(row[4])
		out := row[5] == "true"

		products = append(products, models.Product{
			ID:         0, // assuming upsert
			Name:       row[0],
			Image:      row[1],
			Price:      price,
			Quantity:   qty,
			OutOfStock: out,
		})
	}

	return products, nil
}
