package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintMatrix(Matrix [][]int64, n int, m int) {
	for i := 0; i < n; i++ {
		fmt.Print("[")
		for j := 0; j < m; j++ {
			fmt.Printf("%3d", Matrix[i][j])
		}
		fmt.Println("]")
	}
}

func Multiplicador(A [][]int64, B [][]int64, n int, m int) [][]int64 {

	Matrix := make([][]int64, n)
	for i := 0; i < n; i++ {
		Matrix[i] = make([]int64, m)
	}

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

func Naive(A [][]int64, B [][]int64, C [][]int64, D [][]int64, n int, m int) {

	AB := Multiplicador(A, B, n, m)
	ABC := Multiplicador(AB, C, n, m)
	ABCD := Multiplicador(ABC, D, n, m)

	PrintMatrix(ABCD, n, m)
}

func main() {

	n := 4
	m := 4
	rand.Seed(time.Now().UnixNano())

	A := make([][]int64, n)
	B := make([][]int64, n)
	C := make([][]int64, n)
	D := make([][]int64, n)

	for i := 0; i < n; i++ {
		A[i] = make([]int64, m)
		B[i] = make([]int64, m)
		C[i] = make([]int64, m)
		D[i] = make([]int64, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			A[i][j] = rand.Int63n(10)
			B[i][j] = rand.Int63n(10)
			C[i][j] = rand.Int63n(10)
			D[i][j] = rand.Int63n(10)
		}
	}

	fmt.Println("Matriz A:")
	PrintMatrix(A, n, m)
	fmt.Println("---------")

	fmt.Println("Matriz B:")
	PrintMatrix(B, n, m)
	fmt.Println("---------")

	fmt.Println("Matriz C:")
	PrintMatrix(C, n, m)
	fmt.Println("---------")

	fmt.Println("Matriz D:")
	PrintMatrix(D, n, m)
	fmt.Println("---------")

	Naive(A, B, C, D, n, m)

}
