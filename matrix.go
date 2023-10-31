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

	/* Store the elements in their respective indices, from index [0][0] to [i - 1][j - 1] */
	println("\nInput elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("For matrix[%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	/* Print the matrix */
	fmt.Println("\nMatrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%.2f\t", matrix[i][j])
		}
		fmt.Println()
	}

	/* Print the transposed matrix, inverting the rows and the columns */
	fmt.Println("\nTransposed matrix:")
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			fmt.Printf("%.2f\t", matrix[j][i])
		}
		fmt.Println()
	}

	fmt.Printf("\n%v", matrix) // Print the matrix as a slice
}
