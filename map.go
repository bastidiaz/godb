package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	var key string
	var value int
	var action string

	fmt.Println("This is a simple mapping program. You can store and remove integer values from a map as well as print it.")
	fmt.Println("sv\t- maps a value with a corresponding key; takes in a string key and an integer value (e. g., sv n3 50)")
	fmt.Println("rv\t- deletes a value with and using its corresponding key; takes in a string key (e. g., rv n3)")
	fmt.Println("cm\t- clears the map, deleting all values")
	fmt.Println("pm\t- prints the map")
	fmt.Println("q\t- closes the program")

	for {
		_, err := fmt.Scan(&action)
		if err != nil {
			return
		}

		switch action {
		case "sv":
			_, err := fmt.Scan(&key, &value)
			if err != nil {
				return
			}
			storeValue(myMap, key, value)
			fmt.Println("Value", value, "stored in key", key)
			break
		case "rv":
			_, err := fmt.Scan(&key)
			if err != nil {
				return
			}
			m := myMap[key]
			rmValue(myMap, key)
			fmt.Println("Key", key, "with the value", m, "has been removed")
			break
		case "cm":
			clearMap(myMap)
			break
		case "pm":
			fmt.Println(myMap)
			break
		case "q":
			fmt.Println("Closing program...")
			return
		}
	}
}

func storeValue(myMap map[string]int, key string, value int) {
	myMap[key] = value
}

func rmValue(myMap map[string]int, key string) {
	delete(myMap, key)
}

func clearMap(myMap map[string]int) {
	for key := range myMap {
		delete(myMap, key)
	}
}
