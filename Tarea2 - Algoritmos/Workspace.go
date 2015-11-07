// Tarea 1 - CIT-2001 - Diseño y Análisis de Algoritmos - Universidad Diego Portales
// Integrantes: Guillermo Iglesias Birkner - Josefa González Mejías
// Profesor: Francisco Claude - Ayudante: Marcello Tavano

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
			Matrix[i][j] = (rand.Int63n(1) + 1)
		}
	}
	return Matrix
}

func PrintMatrix(Matrix [][]int64) {
	for i := 0; i < len(Matrix); i++ {
		fmt.Println(Matrix[i])
	}
}

func MatrixGenerator(n, a, b int) [][][]int64 {
	MatrixArray := make([][][]int64, n)
	RNG := make([]int, n+1)

	for i := 0; i < len(RNG); i++ {
		RNG[i] = (rand.Intn(b) + a)
	}

	for i := 0; i < len(MatrixArray); i++ {
		MatrixArray[i] = M_Factory(RNG[i], RNG[i+1])
	}

	return MatrixArray
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
		return
	}

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
		return
	}

	for i := 0; i < (len(M) - 1); i++ {
		aux = int64((len(M[i]) * len(M[i][0]) * len(M[i+1][0])))
		if aux < value {
			value = aux
			save = i
		}
	}

	cost += value

	MxM := Multiplicador(M[save], M[save+1])
	M[save] = MxM

	for i := (save + 1); i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	Greedy(M[:(len(M)-1)], cost)
}

func Top_Down(M [][][]int64) [][]int64 {

	n := (len(M) + 1)

	m := make([][]int64, n)
	s := make([][]int64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
		s[i] = make([]int64, n)
	}

	cost := Top_Down_Func(M, m, s, 1, int64(len(M)))

	NewMatrix := Mult(M, s, 1, int64(len(M)))

	fmt.Println("Costo Total:", cost)

	return NewMatrix
}

func Top_Down_Func(M [][][]int64, m, s [][]int64, i, j int64) int64 {
	if i == j {
		return 0
	}

	m[i][j] = math.MaxInt64

	for k := i; int64(k) <= j-1; k++ {
		q := Top_Down_Func(M, m, s, i, k) + Top_Down_Func(M, m, s, k+1, j) + int64(len(M[i-1])*len(M[k-1][0])*len(M[j-1][0]))
		if q < m[i][j] {
			m[i][j] = q
			s[i][j] = int64(k)
		}
	}

	return m[i][j]
}

func Bottom_Up(M [][][]int64) [][]int64 {
	n := (len(M) + 1)

	m := make([][]int64, n)
	s := make([][]int64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
		s[i] = make([]int64, n)
	}

	for L := 2; L < n; L++ {
		for i := 1; i < n-L+1; i++ {
			j := int64(i + L - 1)
			m[i][j] = math.MaxInt64
			for k := i; int64(k) <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + int64(len(M[i-1])*len(M[k-1][0])*len(M[j-1][0]))
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = int64(k)
				}
			}
		}
	}

	NewMatrix := Mult(M, s, 1, int64(len(M)))

	fmt.Println("Costo Total:", m[1][n-1])
	return NewMatrix
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

	// n, a , b
	Matrices := MatrixGenerator(30, 50, 500)

	var t_prom float64
	var test int = 1

	////////////////////////////////////////////////////////

	fmt.Println("Naive:")
	t_prom = 0
	for i := 0; i < test; i++ {
		Aux := make([][][]int64, len(Matrices))
		copy(Aux, Matrices[0:])
		t1 := time.Now()
		Naive(Aux, 0)
		t2 := time.Since(t1)
		t_prom = t_prom + (float64(t2) / 1000000000)
	}
	t_prom = t_prom / float64(test)
	fmt.Println("Tiempo Ejecución: ", t_prom)
	fmt.Println("---------")

	////////////////////////////////////////////////////////

	fmt.Println("Greedy:")
	t_prom = 0
	for i := 0; i < test; i++ {
		Aux := make([][][]int64, len(Matrices))
		copy(Aux, Matrices[0:])
		t1 := time.Now()
		Greedy(Aux, 0)
		t2 := time.Since(t1)
		t_prom = t_prom + (float64(t2) / 1000000000)
	}
	t_prom = t_prom / float64(test)
	fmt.Println("Tiempo Ejecución: ", t_prom)
	fmt.Println("---------")

	////////////////////////////////////////////////////////

	fmt.Println("Bottom-Up:")
	t_prom = 0
	for i := 0; i < test; i++ {
		Aux := make([][][]int64, len(Matrices))
		copy(Aux, Matrices[0:])
		t1 := time.Now()
		Bottom_Up(Aux)
		t2 := time.Since(t1)
		t_prom = t_prom + (float64(t2) / 1000000000)
	}
	t_prom = t_prom / float64(test)
	fmt.Println("Tiempo Ejecución: ", t_prom)
	fmt.Println("---------")

	////////////////////////////////////////////////////////

	fmt.Println("Top-Down:")
	t_prom = 0
	for i := 0; i < test; i++ {
		Aux := make([][][]int64, len(Matrices))
		copy(Aux, Matrices[0:])
		t1 := time.Now()
		Top_Down(Aux)
		t2 := time.Since(t1)
		t_prom = t_prom + (float64(t2) / 1000000000)
	}
	t_prom = t_prom / float64(test)
	fmt.Println("Tiempo Ejecución: ", t_prom)
	fmt.Println("---------")

}
