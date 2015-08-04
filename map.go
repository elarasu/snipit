package main

import "fmt"

func main() {
	fmt.Println("Testing map")
	var countryList map[string]string
	// make sure its allocated or inited
	countryList = make(map[string]string)
	countryList["us"] = "United States of America"
	countryList["uk"] = "United Kingdom"
	fmt.Println("Lookup us = ", countryList["us"])
	fmt.Println("Lookup uk = ", countryList["uk"])
	fmt.Println("Lookup in = ", countryList["in"])
	// list all values
	for k, v := range countryList {
		fmt.Println("K[", k, "] = ", v)
	}
}
