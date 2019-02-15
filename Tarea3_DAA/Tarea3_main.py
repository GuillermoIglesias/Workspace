# Tarea 3 CIT2001 

import Paris
import Sparta
import Ripley
import re
import sys

# Transformar los numeros en string de una lista a int
def toInt(numero_str):
    digitos = re.findall(r'([0-9]+)', numero_str)
    valor = int(''.join(digitos))
    return valor

# Encontrar el precio minimo para cada talla
def find_Min(Modelo, Tallas, Precio):
	Index_zap = [-1,-1,-1]
	Min_value = [sys.maxint,sys.maxint,sys.maxint]

	for i in range(0,len(Modelo)):
		
		for j in range(0,len(Tallas[i])):
			
			if Tallas[i][j] == str(38):
				p = toInt(Precio[i])
				if p < Min_value[0]:
					Min_value[0] = p
					Index_zap[0] = i
			
			elif Tallas[i][j] == str(40):
				p = toInt(Precio[i])
				if p < Min_value[1]:
					Min_value[1] = p
					Index_zap[1] = i

			elif Tallas[i][j] == str(42):
				p = toInt(Precio[i])
				if p < Min_value[2]:
					Min_value[2] = p
					Index_zap[2] = i

	info_Print(Modelo, Tallas, Precio, Index_zap)

# Imprime resultados
def info_Print(Modelo, Tallas, Precio, Index_zap):
	size = [38,40,42]
	sumPrecio = 0
	for i in range(0,3):
		print 'Marca/Modelo: ' + str(Modelo[Index_zap[i]])
		print 'Talla: ' + str(size[i])
		print 'Precio: ' + str(Precio[Index_zap[i]])
		sumPrecio = sumPrecio + toInt(Precio[Index_zap[i]])
	print '### Valor Total: $' + str(sumPrecio)

def main():
	
	# Modificar estos valores para determinar cuantos productos se mostraran de la pagina
	# Los valores maximos pueden variar en el tiempo
	# Se colocaron con fines de prueba, debido al tiempo de ejecucion con valores muy altos
	Ripley_Size = 100 # Valor maximo 100
	Sparta_Size = 30 # Valor maximo 30
	Paris_Size 	= 200 # Valor maximo 200 

	print 'Cargando informacion Ripley...'
	Modelos_R, Tallas_R, Precios_R = Ripley.Catalogo(Ripley_Size)
	print '\n###################'
	print '## Tienda Ripley ##'
	print '###################\n'
	find_Min(Modelos_R, Tallas_R, Precios_R)

	print '\nCargando informacion Sparta...'
	Modelos_S, Tallas_S, Precios_S = Sparta.Catalogo(Sparta_Size)
	print '\n###################'
	print '## Tienda Sparta ##'
	print '###################\n'
	find_Min(Modelos_S, Tallas_S, Precios_S)

	print '\nCargando informacion Paris...'
	Modelos_P, Tallas_P, Precios_P = Paris.Catalogo(Paris_Size)
	print '\n###################'
	print '## Tienda Paris ##'
	print '###################\n'
	find_Min(Modelos_P, Tallas_P, Precios_P)

if __name__ == "__main__":
    main()
