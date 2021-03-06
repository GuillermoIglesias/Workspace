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
			Matrix[i][j] = rand.Int63n(1) + 1
		}
	}
	return Matrix
}

func PrintMatrix(Matrix [][]int64) {
	for i := 0; i < len(Matrix); i++ {
		fmt.Print("[")
		for j := 0; j < len(Matrix[0]); j++ {
			fmt.Printf("%6d", Matrix[i][j])
		}
		fmt.Println("]")
	}
}

func Multiplicador(A, B [][]int64) [][]int64 {

	Matrix := M_Factory(len(A), len(B[0]))

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[0]); j++ {
			var sum int64 = 0
			for n := 0; n < len(A[0]); n++ {
				sum += (A[i][n] * B[n][j])
			}
			Matrix[i][j] = sum
		}
	}

	return Matrix
}

func Naive(M [][][]int64) {

	if len(M) == 1 {
		fmt.Println("Naive:")
		PrintMatrix(M[0])
		fmt.Println("---------")
		return
	}

	MxM := Multiplicador(M[0], M[1])
	M[0] = MxM

	for i := 1; i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	Naive(M[:(len(M) - 1)])
}

func Greedy(M [][][]int64) {
	var value int = 100000000
	var aux int
	var save int = 0

	if len(M) == 1 {
		fmt.Println("Greedy:")
		PrintMatrix(M[0])
		return
	}

	for i := 0; i < (len(M) - 1); i++ {
		aux = (len(M[i]) * len(M[i][0]) * len(M[i+1][0]))
		if aux < value {
			value = aux
			save = i
		}
	}

	MxM := Multiplicador(M[save], M[save+1])
	M[save] = MxM

	for i := (save + 1); i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	Greedy(M[:(len(M) - 1)])
}

func main() {

	rand.Seed(time.Now().UnixNano())

	A := M_Factory(10, 10)
	B := M_Factory(10, 5)
	C := M_Factory(5, 6)
	D := M_Factory(6, 2)
	E := M_Factory(2, 1)
	F := M_Factory(1, 3)
	G := M_Factory(3, 8)
	H := M_Factory(8, 2)
	I := M_Factory(2, 2)

	Matrices := [][][]int64{A, B, C, D, E, F, G, H, I}

	Aux1 := make([][][]int64, len(Matrices))
	Aux2 := make([][][]int64, len(Matrices))
	copy(Aux1, Matrices[0:])
	copy(Aux2, Matrices[0:])

	fmt.Println("Matriz A:")
	PrintMatrix(A)
	fmt.Println("---------")

	fmt.Println("Matriz B:")
	PrintMatrix(B)
	fmt.Println("---------")

	fmt.Println("Matriz C:")
	PrintMatrix(C)
	fmt.Println("---------")

	fmt.Println("Matriz D:")
	PrintMatrix(D)
	fmt.Println("---------")

	Naive(Aux1)
	Greedy(Aux2)

}
