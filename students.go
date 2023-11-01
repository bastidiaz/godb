package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var studentArr []student

	for {
		var action string
		var idNum int32
		var name string
		var college string

		fmt.Printf("What do you want to do? (add/remove/list/quit): ")
		_, err := fmt.Scanln(&action)
		if err != nil {
			println("Invalid Action")
			continue
		}

		switch action {
		case "add":
			fmt.Printf("Enter ID Number: ")
			_, err := fmt.Scanln(&idNum)
			if err != nil {
				println("Invalid ID Number")
				continue
			} else {
				idNumExists := false
				for _, idExists := range studentArr {
					if idNum == idExists.idNum {
						println("ID Number already exists")
						idNumExists = true
						continue
					}
				}
				if idNumExists {
					continue
				}
			}

			fmt.Printf("Enter Name: ")
			scanner.Scan()
			if scanner.Err() != nil {
				println("Invalid Name")
				continue
			}
			name = scanner.Text()

			fmt.Printf("Enter College: ")
			scanner.Scan()
			if scanner.Err() != nil {
				println("Invalid College")
				continue
			}
			college = scanner.Text()

			addStudent(&studentArr, idNum, name, college)
			break
		case "remove":
			fmt.Printf("Enter ID Number: ")
			_, err := fmt.Scanln(&idNum)
			if err != nil {
				println("Invalid ID Number")
				continue
			}
			removeStudent(&studentArr, idNum)
			break
		case "list":
			for _, student := range studentArr {
				fmt.Printf("%d %s\t%s\n", student.idNum, student.name, student.college)
			}
			break
		case "quit":
			return
		}
	}
}

type student struct {
	idNum   int32
	name    string
	college string
}

func addStudent(studentArr *[]student, idNum int32, name string, college string) {
	*studentArr = append(*studentArr, student{idNum, name, college})
}

func removeStudent(studentArr *[]student, idNum int32) {
	for i, student := range *studentArr {
		if student.idNum == idNum {
			*studentArr = append((*studentArr)[:i], (*studentArr)[i+1:]...)
		}
	}
}
