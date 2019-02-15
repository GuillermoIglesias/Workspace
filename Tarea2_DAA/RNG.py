# Generador de Archivos para Experimentos
# Guarda una lista de elementos de la forma KObjects.txt

import random

def CreateRandom(K):

	new_Random  = open(`K` + "Objects.txt","w")

	# Crea arreglos de K elementos que almacenan nombre, peso y precio
	rnd_name, rnd_weight, rnd_price = [], [], []
	for i in range(K):
		rnd_name.append("item"+str(i+1))
		rnd_weight.append(i%999+1)
		rnd_price.append(i%499+1)
	random.shuffle(rnd_weight)
	random.shuffle(rnd_price)	

	# Se escribe en el archivo de texto los valores aleatorios
	for i in range(K):
		new_Random.write(rnd_name[i])
		new_Random.write(" %s " % rnd_weight[i])
		new_Random.write("%s\n" % rnd_price[i])

	new_Random.close()

def ReadRandom(K):

	read_Random  = open(`K` + "Objects.txt","r")

	name, weight, price = [], [], []
	# Traspasa los textos a dos arreglos
	for line in read_Random:
		values = line.split()
		name.append(values[0])
		weight.append(int(values[1]))
		price.append(int(values[2]))

	read_Random.close()

	return name, weight, price