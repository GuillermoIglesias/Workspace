# Busca todos los resultados de zapatillas de la pagina Sparta.cl
# Retorna 3 arreglos en donde se almacenan los resultados

import urllib2 
import re 

def Sparta():
	# Almacenamiento de todas las zapatillas con sus respectivos datos
	Modelos_Z, Tallas_Z, Precios_Z = [], [], []	

	# URL de la tienda online de Sparta
	psize = str(30)  # Modificar este valor para ver una cantidad maxima de resultados
			       	 # Con un valor de 30 se muestran todos los resultados existentes
			       	 # Se puede reducir este valor para realizar pruebas, debido al tiempo de ejecucion
	url  = 'http://www.sparta.cl/deportes/running/zapatillas-running.html?limit=' + psize
	page = urllib2.urlopen(url)
	html = page.read()

	# Buscar las URL de todas zapatillas
	url_zap = re.findall(r'<a class="area-link" href="http://www.sparta.cl/deportes/running/zapatillas-running/.*html',html)

	# Se revisa el URL de cada zapatilla, debido a que en el URL principal no se encuentran las tallas
	for i in range(0,len(url_zap)):
		# Abrir el URL de cada zapatilla para obtener sus caracteristicas
		page_zap = urllib2.urlopen(url_zap[i][27:])
		html_zap = page_zap.read()

		# Busca el genero 
		genero      = re.findall(r'<td class="data">(.+?)</td>', html_zap)
		genero_str  = re.findall(r'([A-Za-z]+)', genero[0]) # Se limpia el string para dejar solamente el genero

		# Encontrar sus tallas disponibles
		tallas      = re.findall(r',"label":"([^:"a-zA-Z]+)"', html_zap)

		# Algunos productos no tienen tallas disponibles
		if len(tallas) == 0: 
			continue

		# Se comprueba si hay tallas disponibles deseadas 
		tallas_disp = []

		# Se cambian las tallas us por la nacional dependiendo del genero
		if genero_str[0] == 'Hombre':
			for i in range(0,len(tallas)):
				if tallas[i] == str(8) or tallas[i] == str(10):
					tallas_disp.append(tallas[i])

			for i in range(0,len(tallas_disp)):
				if tallas_disp[i] == str(8):
					tallas_disp[i] = str(40)

				if tallas_disp[i] == str(10):
					tallas_disp[i] = str(42)

		if genero_str[0] == 'Mujer':
			for i in range(0,len(tallas)):
				if tallas[i] == str(8) or tallas[i] == str(10):
					tallas_disp.append(tallas[i])

			for i in range(0,len(tallas_disp)):
				if tallas_disp[i] == str(8):
					tallas_disp[i] = str(38)

				if tallas_disp[i] == str(10):
					tallas_disp[i] = str(40)	

		# Sino existen tallas deseadas, no interesa saber mas de esa zapatilla
		if len(tallas_disp) == 0:
			continue

		# Encontrar su modelo y marca
		modelo      = re.findall(r'<h1 class="h1"  itemprop="name">([^<]+)', html_zap)

		# Encontrar su precio
		precio = re.findall(r'<span class="price" id="product-price-....">(.+?)</span>', html, flags = re.DOTALL)
		if len(precio) == 0: # Algunos productos tienen el precio en un formato diferente
			precio = re.findall(r'<span class="price" id="old-price-....">(.+?)</span>', html, flags = re.DOTALL)
		precio_str = re.findall(r'([0-9\.\$]+)', precio[0]) # Se limpia el string para dejar solamente el precio
		if len(precio_str) == 0: # Algunos productos no tienen precio, al no existir stock al momento de consultar
			continue 

		# Almacenamiento de datos correspondientes
		Modelos_Z.append(modelo)
		Tallas_Z.append(tallas_disp)
		Precios_Z.append(precio_str)

		# Solo se utiliza en caso de querer ver las URLs de cada zapatilla elegida
		# print url_zap[i][27:]
		
	return Modelos_Z, Tallas_Z, Precios_Z	