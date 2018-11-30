package main

import "log"

func tryCompose()  {
	var product = Product{
		Price: 90000000,
		Stock: 9900,
		Name: "product one",
		ID: 1,
	}
	
	var validateProduct = ValidateName.
		Compose(ValidatePrice).
		Compose(ValidateStock)
	var isProductValid, err = validateProduct(product)
	log.Println("tryValidateProduct", isProductValid, err)

	// what if there is a new business process
	// and the validation is just slightly different?
	// are we just gonna make a whole new code for it?
	// no, simply reuse the typedef, add new func on top of it
	var validateProductV3 = ValidateNameV3.
		Compose(ValidatePrice).
		Compose(ValidateStock)
	isProductValid, err = validateProductV3(product)
	log.Println("tryValidateProductV3", isProductValid, err)

}
