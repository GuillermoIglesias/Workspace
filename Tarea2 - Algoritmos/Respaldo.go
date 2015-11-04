package main

import (
	"fmt"
	"math/rand"
	"time"
)

func M_Factory(n, m int) [][]int64 {
	Matrix := make([][]int64, n)
	for i := 0; i < n; i++ {
		Matrix[i] = make([]int64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Matrix[i][j] = 1
		}
	}
	return Matrix
}

func PrintMatrix(Matrix [][]int64, n, m int) {
	for i := 0; i < n; i++ {
		fmt.Print("[")
		for j := 0; j < m; j++ {
			fmt.Printf("%3d", Matrix[i][j])
		}
		fmt.Println("]")
	}
}

func Multiplicador(A, B [][]int64, n, m int) [][]int64 {

	Matrix := M_Factory(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var sum int64 = 0
			for a := 0; a < n; a++ {
				sum += (A[i][a] * B[a][j])
			}
			Matrix[i][j] = sum
		}
	}

	return Matrix
}

func Naive(A, B, C, D [][]int64) {

	AB := Multiplicador(A, B, len(A), len(B[0]))
	ABC := Multiplicador(AB, C, len(AB), len(C[0]))
	ABCD := Multiplicador(ABC, D, len(ABC), len(D[0]))

	PrintMatrix(ABCD, len(ABCD), len(ABCD[0]))
}

func Greedy(A, B, C, D [][]int64) {

}

func main() {

	rand.Seed(time.Now().UnixNano())

	A := M_Factory(2, 3)
	B := M_Factory(3, 4)
	C := M_Factory(4, 3)
	D := M_Factory(3, 2)

	fmt.Println("Matriz A:")
	PrintMatrix(A, len(A), len(A[0]))
	fmt.Println("---------")

	fmt.Println("Matriz B:")
	PrintMatrix(B, len(B), len(B[0]))
	fmt.Println("---------")

	fmt.Println("Matriz C:")
	PrintMatrix(C, len(C), len(C[0]))
	fmt.Println("---------")

	fmt.Println("Matriz D:")
	PrintMatrix(D, len(D), len(D[0]))
	fmt.Println("---------")

	Naive(A, B, C, D)

}
