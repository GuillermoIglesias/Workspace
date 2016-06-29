# Busca todos los resultados de zapatillas de la pagina Paris.cl
# Retorna 3 arreglos en donde se almacenan los resultados

import urllib2 
import re 

def Paris():
	# Almacenamiento de todas las zapatillas con sus respectivos datos
	Modelos_Z, Tallas_Z, Precios_Z = [], [], []

	# URL de la tienda online de Paris
	mcan = str(400) # Modificar este valor para ver una cantidad maxima de resultados
			       	# Con un valor de 400 se muestran todos los resultados existentes
			       	# Se puede reducir este valor para realizar pruebas, debido al tiempo de ejecucicion
	url  = 'http://www.paris.cl/webapp/wcs/stores/servlet/SearchDisplay?searchTermScope=&searchType=1000&filterTerm=&orderBy=&maxPrice=&showResultsPage=true&langId=-5&beginIndex=0&sType=SimpleSearch&metaData=YWRzX2YyNTUwMV9udGtfY3M6IlJ1bm5pbmci&manufacturer=&resultCatEntryType=&catalogId=40000000629&pageView=image&searchTerm=&minPrice=&categoryId=51386230&storeId=10801&pageSize=' + mcan
	page = urllib2.urlopen(url)
	html = page.read()

	# Buscar las URL de todas zapatillas
	url_zap = re.findall(r'<a  href="http://www.paris.cl/tienda/es/paris/moda-calzado/moda-calzado/.*-',html)

	for i in range(0,len(url_zap)):
		# Abrir el URL de cada zapatilla para obtener sus caracteristicas 
		page_zap = urllib2.urlopen(url_zap[i][10:])
		html_zap = page_zap.read()

		# Encontrar sus tallas disponibles
		tallas = re.findall(r'Talla_CL ([^ _|:"a-zA-Z]+)', html_zap)
		# Algunos productos no tienen tallas disponibles
		if len(tallas) == 0: 
			continue
		
		# Se comprueba si hay tallas disponibles deseadas
		tallas_disp = []
		for i in range(0,len(tallas)):
			if tallas[i] == str(38) or tallas[i] == str(40) or tallas[i] == str(42):
				tallas_disp.append(tallas[i])
		# Sino existen tallas deseadas, no interesa saber mas de esa zapatilla
		if len(tallas_disp) == 0:
			continue

		# Encontrar su modelo y marca
		modelo = re.findall(r'<meta property="og:description" content="([^<]+)"', html_zap)		

		# Encontrar su precio
		precio = re.findall(r'<div id="price" class="price offerPrice bold">(.+?)</div>', html_zap, flags = re.DOTALL)
		if len(precio) == 0: # Algunos productos tienen el precio en un formato diferente
			precio = re.findall(r'<div id="offerPrice" class="price offerPrice bold">(.+?)</div>', html_zap, flags = re.DOTALL)
		precio_str = re.findall(r'([0-9\.\$]+)', precio[0]) # Se limpia el string para dejar solamente el precio
		if len(precio_str) == 0: # Algunos productos no tienen precio, al no existir stock al momento de consultar
			continue 
		
		# Almacenamiento de datos correspondientes
		Modelos_Z.append(modelo)
		Tallas_Z.append(tallas_disp)
		Precios_Z.append(precio_str)

		# Solo se utiliza en caso de querer ver las URLs de cada zapatilla elegida
		# print url_zap[i][10:]
		
	return Modelos_Z, Tallas_Z, Precios_Z	