package main

import "log"

const imperativePrefix = "[Imperative]"

func tryImperativeAny()  {
	// check if productID 1 is exist
	var productIDOne int64 = 1
	var isProductOneExist bool
	for _, product := range productSlice {
		if product.ID == productIDOne {
			isProductOneExist = true
			break
		}
	}
	log.Println(imperativePrefix, "isProductOneExist", isProductOneExist)

	// what if we want to check another condition?
	// we must loop again with the same bloated code
	var productIDTwo int64 = 2
	var isProductTwoExist bool
	for _, product := range productSlice {
		if product.ID == productIDTwo {
			isProductTwoExist = true
			break
		}
	}
	log.Println(imperativePrefix, "isProductTwoExist", isProductTwoExist, "sadly 7 line!!")
}

func tryImperativeFilter() {
	// get product with price greater than 40k and stock greater than 50
	// first we have to declare the container
	var priceToFilter float64 = 40000
	var stockToFilter int64 = 50
	var productSliceOne ProductSlice
	for _, product := range productSlice {
		if product.Price > priceToFilter && product.Stock > stockToFilter {
			productSliceOne = append(productSliceOne, product)
		}
	}
	log.Println(imperativePrefix, "productSliceOne", productSliceOne)

	// what if we want another condition?
	// again, we have to repeat writing the same loop code!
	var idToFilterNot int64 = 9
	// first we have to declare the container
	var productSliceTwo ProductSlice
	for _, product := range productSlice {
		if product.Price > priceToFilter && product.Stock > stockToFilter && product.ID != idToFilterNot {
			productSliceTwo = append(productSliceTwo, product)
		}
	}
	log.Println(imperativePrefix, "productSliceTwo", productSliceTwo)
}