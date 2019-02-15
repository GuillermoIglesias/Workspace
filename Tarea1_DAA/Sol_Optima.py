# Solucion Optima 
# Entrega el kth termino de un arreglo
# Y los Top M menores numeros de un arreglo
# Mejor y peor caso en ambos O(n)

import heapq

# Mediante esta funcion, es posible para el algoritmo de Seleccion, 
# encontrar un pivote optimo para resolver el problema,
# entregando asi los indices correspondientes de cada particion.
def Partition(A, left, right, pivot): 
	index_swap = left 

	for i in range(left, right+1): 
		if A[i] < pivot: 
			A[i], A[index_swap] = A[index_swap], A[i] 
			index_swap += 1 
	
	return index_swap - 1   

# Modificacion de QuickSelect que permite encontrar un pivote optimo,
# mediante el uso de la mediana de las medianas.
def Median_of_Medians(A, left, right, k): 
	length = right-left+1 
	
	# Define el largo maximo de cada particion
	if length <= 5: 
		return sorted(A[left:right+1])[k-1]   
	
	numMedians	= length/5 
	medians 	= [Median_of_Medians(A, left+5*i, left+5*(i+1)-1, 3) for i in range(numMedians)]
	pivot 		= Median_of_Medians(medians, 0, len(medians)-1, len(medians)/2+1) 
	index_pivot = Partition(A, left, right, pivot) 
	rank 		= index_pivot-left+1 
	
	if k <= rank: 
		return Median_of_Medians(A, left, index_pivot, k) 
	else: 
		return Median_of_Medians(A, index_pivot+1, right, k-rank) 


def Min_Heap(A, m):
	# Crea un heap a partir de una lista en O(n)
	heapq.heapify(A)
	
	# Almacena el menor elemento en un arreglo
	# Cada operacion se realiza en O(logn)
	top_m = []
	for i in range(m):
		top_m.append(heapq.heappop(A))
	return top_m

def kth_optimo(A,k):
	return Median_of_Medians(A,0,len(A)-1,k)

def top_m_optimo(A,m):
	# Se utiliza una copia de la lista para no alterar la original
	B = A[:]
	return Min_Heap(B,m)