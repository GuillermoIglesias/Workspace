# Solucion Basica 
# Entrega el kth termino de un arreglo
# Y los Top M menores numeros de un arreglo
# Mejor y peor caso en ambos es O(nlogn)

import heapq

# Ordena un arreglo mediande Divide y Conquista
def MergeSort(A):
	# Caso Base
	if len(A) <= 1 :
		return A

	# Separacion del arreglo en dos mitades
	mid 	= len(A) / 2
	left 	= A[:mid]
	right 	= A[mid:]
	
	# Divide el arreglo recursivamente
	left 	= MergeSort(left)
	right 	= MergeSort(right)

	# Combina los resultados en tiempo lineal O(n)
	return list(heapq.merge(left,right)) 

def kth_sencillo(A,k):
	B = MergeSort(A)
	return B[k-1]

def top_m_sencillo(A,m):
	B = MergeSort(A)
	return B[:m]