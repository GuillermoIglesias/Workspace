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

func PrintMatrix(Matrix [][]int64) {
	for i := 0; i < len(Matrix); i++ {
		fmt.Print("[")
		for j := 0; j < len(Matrix[0]); j++ {
			fmt.Printf("%3d", Matrix[i][j])
		}
		fmt.Println("]")
	}
}

func Multiplicador(A, B [][]int64) [][]int64 {

	Matrix := M_Factory(len(A), len(B[0]))

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[0]); j++ {
			var sum int64 = 0
			for a := 0; a < len(A); a++ {
				sum += (A[i][a] * B[a][j])
			}
			Matrix[i][j] = sum
		}
	}

	return Matrix
}

func Naive(M [][][]int64) {

	AB := Multiplicador(M[0], M[1])
	ABC := Multiplicador(AB, M[2])
	ABCD := Multiplicador(ABC, M[3])

	fmt.Println("Naive:")
	PrintMatrix(ABCD)
	fmt.Println("---------")
}

func Greedy(M [][][]int64) {
	var value int = 10000
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
		M[i+1] = nil
	}

	Greedy(M[:(len(M) - 1)])
	return
}

func main() {

	rand.Seed(time.Now().UnixNano())

	A := M_Factory(2, 3)
	B := M_Factory(3, 4)
	C := M_Factory(4, 3)
	D := M_Factory(3, 2)

	Matrices := [][][]int64{A, B, C, D}

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

	Naive(Matrices)

	Greedy(Matrices)

}
