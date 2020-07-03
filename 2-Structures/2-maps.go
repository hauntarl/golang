package main

import "fmt"

func main() {
	/*
		Maps:
			- key can be anything on which we can perform equality operation
			- eg. booleans, numeric types, strings, pointers, interfaces, structs, arrays, channels
			- slices, maps and functions are not allowed
			- order of elements inserted is not maintained
			- value of uninitialized map is nil
			- maps are pass by reference like slices (shallow copy)

		Map Creation:
			1. map[<key-datatype>]<value-datatype>{<pairs...>}
				- <key-datatype> should match above constraints
				- <pairs...> are the entries in the map
				- use this if you already know what data needs to be inserted in the map
			2. using built-in make() func
				- make(map[<key-datatype>]<value-datatype>, <capacity>)
				- capacity parameter is optional, it helps compiler allocate
				  memory accordingly
				- use this if you have rough idea about how many elements you'll be
				  accomodating in the map, irrespective of if you mention it, map will still
				  grow, the only difference it makes is it helps you set the initial size of map
				- The initial capacity does not bound its size: maps grow to accommodate
				  the number of items stored in them, with the exception of nil maps.
				- A nil map is equivalent to an empty map except that no elements may be added.


		Deleting from Map:
			- built-in delete() function
			- delete(<name-of-map>, <key>)

		Retrieving Data:
			- <map>[<key>]
			- if key absent it returns default value of <value-datatype>
			- which may cause some problems, but retrieving returns 2 variables
			- data, ok := myMap["key"]
			- ok is a bool value which tells us whether the key exists in map or not
			- it is a convention to use "ok" as variable name
			- if you are just interested in whether key exists or not,
			  use "_" write-only operator to throw away the data

		Operations:
			- find length using len() function
	*/
	fmt.Println("\nMAP CREATION:\nWhen values are known")
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	fmt.Println(statePopulation)

	fmt.Println("\nWhen values are unknown")
	countryPopulation := make(map[string]int)
	countryPopulation["India"] = 1352600000 // adding a new pair in map
	fmt.Println("Country Population India: ", countryPopulation["India"])

	fmt.Println("\nMAP DELETION:")
	data, ok := statePopulation["New York"]
	fmt.Println("Before deleting: New York pop.: ", data, "\n\tIsPresent: ", ok)
	fmt.Println("\tmap length: ", len(statePopulation))
	delete(statePopulation, "New York")

	data, ok = statePopulation["New York"]
	fmt.Println("After deleting: New York pop.: ", data, "\n\tIsPresent: ", ok)
	fmt.Println("\tmap length: ", len(statePopulation))

	_, ok = statePopulation["New York"]
	fmt.Println("Is New York data present?", ok)
}
