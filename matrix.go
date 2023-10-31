package main

import (
	"fmt"
)

func main() {
	var rows, columns int

	println("Enter the number of rows and columns [e. g., r c]: ")
	_, err := fmt.Scan(&rows, &columns)
	if err != nil {
		return
	}

	matrix := make([][]float32, rows)

	for i := range matrix {
		matrix[i] = make([]float32, columns)
	}

	println("Input elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("For matrix[%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("Matrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%.2f\t", matrix[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Transposed matrix:")
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			fmt.Printf("%.2f\t", matrix[j][i])
		}
		fmt.Println()
	}
}
