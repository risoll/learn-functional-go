package main

import (
	"log"
	"errors"
)

const validateProductPrefix = "[ValidateProduct]"

type ValidateProduct func(product Product) (isValid bool, err error)

func (f ValidateProduct) Compose(innerFunc ValidateProduct) ValidateProduct {
	return func(product Product) (isValid bool, err error) {
		isValid, err = innerFunc(product)
		if err != nil {
			return isValid, err
		}
		if !isValid {
			return isValid, err
		}
		return f(product)
	}
}

const maxProductNameLength = 10
var ValidateName ValidateProduct = func(product Product) (isValid bool, err error) {
	log.Println("name")
	if len(product.Name) > maxProductNameLength {
		return false, errors.New("invalid name")
	}
	return true, nil
}

const maxProductNameLengthV3 = 20
var ValidateNameV3 ValidateProduct = func(product Product) (isValid bool, err error) {
	log.Println("namev3")
	if len(product.Name) > maxProductNameLengthV3 {
		return false, errors.New("invalid name V3")
	}
	return true, nil
}

const maxProductStock = 10000
var ValidateStock ValidateProduct = func(product Product) (isValid bool, err error) {
	log.Println("stock")
	if product.Stock > maxProductStock {
		return false, errors.New("invalid stock")
	}
	return true, nil
}

const maxProductPrice = 100000000
var ValidatePrice ValidateProduct = func(product Product) (isValid bool, err error) {
	log.Println("price")
	if product.Price > maxProductPrice {
		return false, errors.New("invalid price")
	}
	return true, nil
}