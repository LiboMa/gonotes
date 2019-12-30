package main

import "fmt"

func recursion(x int) {

	if x >= 0 {
		x = x - 1
		fmt.Println("recursion decrease ", x)
		recursion(x)
	} else {

		fmt.Println("recursion done ", x)

	}
}

func main() {

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	var x int = 10000

	/* map key value for capital*/

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"
	countryCapitalMap["China"] = "Beijing"

	for country := range countryCapitalMap {
		fmt.Println(country)
	}

	recursion(x)
}
