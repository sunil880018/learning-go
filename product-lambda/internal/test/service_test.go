package test

import (
	"product-lambda/internal/utils"
	"testing"
)

func TestParseCSV(t *testing.T) {
	products, err := utils.ParseCSV("sample_products.csv")
	if err != nil {
		t.Fatal(err)
	}

	if len(products) == 0 {
		t.Fatal("No products parsed")
	}
}
