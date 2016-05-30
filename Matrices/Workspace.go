/*
	Tarea 2 - CIT-2001 - Diseño y Análisis de Algoritmos - Universidad Diego Portales
   	Integrantes: Guillermo Iglesias Birkner - Josefa González Mejías
   	Profesor: Francisco Claude - Ayudante: Marcello Tavano
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Retorna Matriz dimensión n x m.
func M_Factory(n, m int) [][]int64 {
	// Declaración nueva Matriz.
	Matrix := make([][]int64, n)
	for i := 0; i < n; i++ {
		Matrix[i] = make([]int64, m)
	}

	// Relleno de la Matriz, valores random son configurables.
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Matrix[i][j] = rand.Int63n(3)
		}
	}

	return Matrix
}

// Imprime una Matriz, se utilizó para comprobaciones.
func PrintMatrix(Matrix [][]int64) {
	for i := 0; i < len(Matrix); i++ {
		fmt.Println(Matrix[i])
	}
}

// Retorna "arreglo" de n Matrices.
func MatrixGenerator(n, a, b int) [][][]int64 {
	// Declaración variables a utilizar.
	MatrixArray := make([][][]int64, n)
	RNG := make([]int, n+1)

	// Se rellena el arreglo con valores random de rango [a,b].
	for i := 0; i < len(RNG); i++ {
		RNG[i] = (rand.Intn(b) + a)
	}

	// Se construyen matrices a partir del arreglo RNG.
	// RNG es de la forma [a1, a2, a3...an].
	// Luego las matrices son de dimension [ai][ai+1]
	for i := 0; i < len(MatrixArray); i++ {
		MatrixArray[i] = M_Factory(RNG[i], RNG[i+1])
	}

	return MatrixArray
}

// Retorna una Matriz resultante de multiplicacion de 2 Matrices.
func Multiplicador(A, B [][]int64) [][]int64 {
	// Nueva matriz donde se almacenará el resultado.
	Matrix := M_Factory(len(A), len(B[0]))

	//Algoritmo que realizará la multiplicacion de las matrices.
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

// Algoritmo Naive (recursivo), multiplica las matrices en el orden original.
// Por ej. (((((A)B)C)D)E).
func Naive(M [][][]int64, cost int64) {
	// Condición de salida de la recursividad.
	if len(M) == 1 {
		fmt.Println("Costo Total:", cost)
		return
	}

	// El costo de cada paso de 2 matrices NxM * MxK es N*M*K
	cost += int64((len(M[0]) * len(M[0][0]) * len(M[1][0])))

	// Se agrupan las dos primeras matrices multiplicadas entre si.
	MxM := Multiplicador(M[0], M[1])
	M[0] = MxM

	// La segunda matriz se desplaza hacia el fondo del arreglo.
	for i := 1; i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	// Se vuelve a llamar recursivamente a la función, con un valor menos.
	// Este valor se encuentra agregado mediante la multiplicación.
	// Por ej. (((AB)C)D) -> ((ABC)D)
	Naive(M[:(len(M)-1)], cost)
}

// Algoritmo Greedy (recursivo), selecciona que matrices multiplicar según el menor costo sus valores.
// Funciona de forma análoga a la función Naive.
func Greedy(M [][][]int64, cost int64) {
	// Inicialización de valores a utilizar.
	var value int64 = math.MaxInt64
	var aux int64
	var save int = 0

	// Condición de salida de la recursividad.
	if len(M) == 1 {
		fmt.Println("Costo Total:", cost)
		return
	}

	// Se evaluan todos los costos de cada multiplicación.
	// Se utiliza como criterio de costo la cantidad de multiplicaciones.
	for i := 0; i < (len(M) - 1); i++ {
		aux = int64((len(M[i]) * len(M[i][0]) * len(M[i+1][0])))
		if aux < value {
			value = aux
			save = i
		}
	}

	// El costo de cada paso de 2 matrices NxM * MxK es N*M*K
	cost += value

	// Se agrupan las dos matrices elegidas multiplicadas entre si.
	MxM := Multiplicador(M[save], M[save+1])
	M[save] = MxM

	// La matriz siguiente a la elegida se desplaza hacia el fondo del arreglo.
	for i := (save + 1); i < (len(M) - 1); i++ {
		M[i] = M[i+1]
	}

	// Se vuelve a llamar recursivamente a la función, con un valor menos.
	// Este valor se encuentra agregado mediante la multiplicación.
	Greedy(M[:(len(M)-1)], cost)
}

// Llamada de la función Top-Down, se utilizó para facilitar su manipulación.
func Top_Down(M [][][]int64) [][]int64 {
	// Inicialización de variables.
	// m[i][j] se utiliza para medir el costo de todas las multiplicaciones.
	// s[i][j] guarda el orden de los paréntesis para su posterior multiplicación.
	n := (len(M) + 1)
	m := make([][]int64, n)
	s := make([][]int64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
		s[i] = make([]int64, n)
	}

	// Costo total de cantidad de operaciones a realizar.
	cost := Top_Down_Func(M, m, s, 1, int64(len(M)))

	// Esta matriz almacena el resultado obtenido.
	NewMatrix := Mult(M, s, 1, int64(len(M)))

	fmt.Println("Costo Total:", cost)
	return NewMatrix
}

// Esta funcion  ejecuta el algoritmo recursivo Top-Down.
// Retorna el costo total de operaciones, modifica m[i][j] y s[i][j].
func Top_Down_Func(M [][][]int64, m, s [][]int64, i, j int64) int64 {
	// Para casos donde se considera la misma matriz, será 0.
	if i == j {
		return 0
	}

	// Inicialización de la variable.
	m[i][j] = math.MaxInt64

	// Intenta todos los posibles casos.
	for k := i; int64(k) <= j-1; k++ {
		// q es el costo de hacer todas las multiplicaciones.
		q := Top_Down_Func(M, m, s, i, k) + Top_Down_Func(M, m, s, k+1, j) + int64(len(M[i-1])*len(M[k-1][0])*len(M[j-1][0]))
		// Cuando el valor es menor, guarda los resultados en las variables correspondientes.
		if q < m[i][j] {
			m[i][j] = q
			s[i][j] = int64(k)
		}
	}

	return m[i][j]
}

// Algoritmo Bottom-Up, toma la secuencia de matrices y las separa en dos subsecuencias.
// Encuentra el mínimo costo de multiplicar cada subsecuencia.
func Bottom_Up(M [][][]int64) [][]int64 {
	// Inicialización de variables.
	// m[i][j] se utiliza para medir el costo de todas las multiplicaciones.
	// s[i][j] guarda el orden de los paréntesis para su posterior multiplicación.
	n := (len(M) + 1)
	m := make([][]int64, n)
	s := make([][]int64, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
		s[i] = make([]int64, n)
	}

	// L es el largo de la cadena.
	for L := 2; L < n; L++ {
		for i := 1; i < n-L+1; i++ {
			j := int64(i + L - 1)
			// Inicialización de la variable.
			m[i][j] = math.MaxInt64
			for k := i; int64(k) <= j-1; k++ {
				// q es el costo de hacer todas las multiplicaciones.
				q := m[i][k] + m[k+1][j] + int64(len(M[i-1])*len(M[k-1][0])*len(M[j-1][0]))

				// Cuando el valor es menor, guarda los resultados en las variables correspondientes.
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = int64(k)
				}
			}
		}
	}

	// Esta matriz almacena el resultado obtenido.
	NewMatrix := Mult(M, s, 1, int64(len(M)))

	fmt.Println("Costo Total:", m[1][n-1])
	return NewMatrix
}

// Una vez encontradas las posiciones de los parentesis y almacenadas en s[i][j].
// Mult procede a realizar las multiplicaciones en dicho orden.
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
	// Seed utilizada para números random.
	rand.Seed(time.Now().UnixNano())

	// Genera n Matrices de dimensión entre [a,b]
	// MatrixGenerator(n,a,b)
	Matrices := MatrixGenerator(10, 10, 1000)

	var t_prom float64
	var test int = 10 //Cantidad de tests para promediarlos.

	/*	Para la realización de este experimento, se utilizan X tests,
		para cada test se obtiene un valor de tiempo, el cual será promediado.
		Cada test utiliza una copia del arreglo de matrices generado.
		Cada apartado entrega el costo (cantidad escalares de multiplicaciones),
		y su tiempo de ejecución en segundos, de cada Algoritmo.  */

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

	////////////////////////////////////////////////////////

}
