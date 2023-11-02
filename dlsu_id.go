package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("A simple CLI program for verifying if an ID number fits the DLSU format. Only tested with ID 100 and above.")
	fmt.Println("You may select verify (include ID number to check; 'verify 11722339'), generate, or quit.")
	for {
		var action string
		_, err := fmt.Scan(&action)
		if err != nil {
			fmt.Println("Invalid command!")
			continue
		}

		switch action {
		case "verify":
			var idNum int
			_, err := fmt.Scanln(&idNum)
			if err != nil {
				return
			}
			fmt.Printf("The ID number %d is %t\n", idNum, verify(idNum))
			break
		case "generate":
			fmt.Println(generate())
			break
		case "quit":
			return
		}
	}
}

func verify(idNum int) bool {
	verified := false
	var total int

	for i := 1; i < 9; i++ {
		temp := idNum % 10
		idNum /= 10
		total += temp * i
	}

	if total%11 == 0 {
		verified = true
	}

	return verified
}

func generate() int {
	rand.NewSource(time.Now().UnixNano())

	for {
		randomID := rand.Intn(9999999) + 10000000
		if verify(randomID) {
			return randomID
		}
	}
}
