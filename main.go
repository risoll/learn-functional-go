package main

import "log"

var productSlice ProductSlice

func main() {

	// init data
	productSlice = initData()

	//log.Println("try Any/IsExist Func")
	//tryDeclarative()
	//tryImperativeAny()
	//log.Println()
	//
	//log.Println("try Filter/Select Func")
	//tryDeclarativeFilter()
	//tryImperativeFilter()
	//log.Println()

	log.Println("try compose func")
	tryCompose()
}