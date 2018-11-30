package main

import "log"

const declarativePrefix = "[Declarative]"

func tryDeclarative()  {
	// check if productID 1 is exist
	var productIDOne int64 = 1
	var isProductOneExist = productSlice.Any(func(product Product) bool {
		return product.ID == productIDOne
	})
	log.Println(declarativePrefix, "isProductOneExist", isProductOneExist)

	// what if we want to check another condition?
	// we just have to reuse the Any lambda function,
	// pass the desired condition as a function
	var productIDTwo int64 = 1
	var isProductTwoExist = productSlice.Any(func(product Product) bool {
		return product.ID == productIDTwo
	})
	log.Println(declarativePrefix, "isProductTwoExist", isProductTwoExist, "only 3 line!!")

	// let's try more complex condition
	var price40k float64 = 40000
	// normally, we could just directly pass the function to the Any func
	// this is just for the sake of clarity
	var checkComplexFunc = func(product Product) bool {
		return (product.ID == productIDOne ||
			product.ID == productIDTwo) &&
			product.Price >= price40k
	}
	var isCheckComplexProductExist = productSlice.Any(checkComplexFunc)
	log.Println(declarativePrefix, "isCheckComplexProductExist", isCheckComplexProductExist)
}

func tryDeclarativeFilter()  {
	var priceToFilter float64 = 40000
	var stockToFilter int64 = 50
	var productSliceOne = productSlice.Where(func(product Product) bool {
		return product.Price > priceToFilter && product.Stock > stockToFilter
	})
	log.Println(declarativePrefix, "productSliceOne", productSliceOne)

	// what if we want to add another condition?
	// just reuse the lambda function
	var idToFilterNot int64 = 9
	var productSliceTwo = productSlice.Where(func(product Product) bool {
		return product.Price > priceToFilter && product.Stock > stockToFilter && product.ID != idToFilterNot
	})
	log.Println(declarativePrefix, "productSliceTwo", productSliceTwo)
}
