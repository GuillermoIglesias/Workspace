package main

import (
	"fmt"
	"math"
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
			Matrix[i][j] = rand.Int63n(10)
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

func Naive(M [][][]int64, cost int64) {

	if len(M) == 1 {
		fmt.Println("Costo Total:", cost)
		//PrintMatrix(M[0])
		return
	}

	fmt.Println("Value ->", (len(M[0]) * len(M[0][0]) * len(M[1][0])))
	cost += int64((len(M[0]) * len(M[0][0]) * len(M[1][0])))

	MxM := Multiplicador(M[0], M[1])
	M[0] = MxM

	for i := 1; i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	Naive(M[:(len(M)-1)], cost)
}

func Greedy(M [][][]int64, cost int64) {
	var value int64 = math.MaxInt64
	var aux int64
	var save int = 0

	if len(M) == 1 {
		fmt.Println("Costo Total:", cost)
		//PrintMatrix(M[0])
		return
	}

	for i := 0; i < (len(M) - 1); i++ {
		aux = int64((len(M[i]) * len(M[i][0]) * len(M[i+1][0])))
		if aux < value {
			value = aux
			save = i
		}
	}

	fmt.Println("Value ->", value)

	cost += value

	MxM := Multiplicador(M[save], M[save+1])
	M[save] = MxM

	for i := (save + 1); i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	Greedy(M[:(len(M)-1)], cost)
}

func MatrixChainOrder(M [][][]int64) int64 {
	/* For simplicity of the program, one extra row and one extra column are
	allocated in m[][].  0th row and 0th column of m[][] are not used */
	n := (len(M) + 1)

	m := make([][]int64, n)
	s := make([][]int64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
		s[i] = make([]int64, n)
	}

	/* m[i,j] = Minimum number of scalar multiplications needed to compute
	   the matrix A[i]A[i+1]...A[j] = A[i..j] where dimention of A[i] is
	   p[i-1] x p[i] */
	// cost is zero when multiplying one matrix.

	// L is chain length.
	for L := 2; L < n; L++ {
		for i := 1; i < n-L+1; i++ {
			j := int64(i + L - 1)
			m[i][j] = math.MaxInt64
			for k := i; int64(k) <= j-1; k++ {
				// q = cost/scalar multiplications
				q := m[i][k] + m[k+1][j] + int64(len(M[i-1])*len(M[k-1][0])*len(M[j-1][0]))
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = int64(k)
				}
			}
		}
	}

	//PrintMatrix(Mult(M, s, 1, int64(n-1)))
	Mult(M, s, 1, int64(n-1))
	return m[1][n-1]
}

func Mult(M [][][]int64, s [][]int64, i, j int64) [][]int64 {
	if i == j {
		return M[i-1]
	} else {
		var k int64 = s[i][j]
		A := Mult(M, s, i, k)
		B := Mult(M, s, k+1, j)
		return Multiplicador(A, B)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	A := M_Factory(3000, 350)
	B := M_Factory(350, 1500)
	C := M_Factory(1500, 500)
	D := M_Factory(500, 1000)
	E := M_Factory(1000, 2000)
	F := M_Factory(2000, 2500)

	Matrices := [][][]int64{A, B, C, D, E, F}

	Aux1 := make([][][]int64, len(Matrices))
	Aux2 := make([][][]int64, len(Matrices))
	Aux3 := make([][][]int64, len(Matrices))
	copy(Aux1, Matrices[0:])
	copy(Aux2, Matrices[0:])
	copy(Aux3, Matrices[0:])

	fmt.Println("Naive:")
	Naive(Aux1, 0)
	fmt.Println("---------")

	fmt.Println("Greedy:")
	Greedy(Aux2, 0)
	fmt.Println("---------")

	fmt.Println("Matrix Chain Order:")
	fmt.Println("Costo Minimo:", MatrixChainOrder(Aux3))
	fmt.Println("---------")
}
