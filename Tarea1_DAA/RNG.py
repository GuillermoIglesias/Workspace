# Generador de Archivos para Experimentos
# Guarda una lista de elementos de la forma type_numberMax.txt

import random

def CreateRandom(n):

	new_Random  = open("Random_" + `n` + ".txt","w")

	# Crea un arreglo de n elementos en orden
	Random = []
	for i in range(n):
		Random.append(i+1)
		# Aleatoriza la lista generada en orden	
	random.shuffle(Random)

	# Se escribe en el archivo de texto los valores aleatorios
	for i in range(n):
		new_Random.write("%s\n" % Random[i])

	new_Random.close()

def ReadRandom(n):

	read_Random  = open("Random_" + `n` + ".txt","r")

	Random = []
	# Traspasa los textos a dos arreglos
	for line in read_Random:
		Random.append(int(line))

	read_Random.close()

	return Random

def CreateOrden(n):

	Orden = []
	for i in range(n):
		Orden.append(i+1)

	return Orden