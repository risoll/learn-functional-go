package main

import "fmt"

func initData() ProductSlice  {
	var productSlice ProductSlice
	const basePrice = 10000
	const baseName = "productName"
	const baseStock = 10
	for _, i := range []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		productSlice = append(productSlice, Product{
			ID: i,
			Price: float64(basePrice * i),
			Name: fmt.Sprintf("%s-%d", baseName, i),
			Stock: baseStock * i,
		})
	}

	return productSlice
}
