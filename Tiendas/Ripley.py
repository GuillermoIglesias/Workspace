# Busca todos los resultados de zapatillas de la pagina Ripley.cl
# Retorna 3 arreglos en donde se almacenan los resultados

import urllib2 
import re 

def Ripley():
	# Almacenamiento de todas las zapatillas con sus respectivos datos
	Modelos_Z, Tallas_Z, Precios_Z = [], [], []

	# URL de la tienda online de Paris
	psize = str(50) # Modificar este valor para ver una cantidad maxima de resultados
			       	 # Con un valor de 100 se muestran todos los resultados existentes
			       	 # Se puede reducir este valor para realizar pruebas, debido al tiempo de ejecucion
	url  	= 'http://simple.ripley.cl/deporte/zapatillas?pageSize=' + psize + '&orderBy=3&facet=ads_f712001_ntk_cs%253A%2522Running%2522'
	# En el caso exclusivo de Ripley, se necesito agregar headers para que la pagina ignorara que es un script maligno
	headers = {'User-Agent' : 'Mozilla/5.0 AppleWebKit/537.36 Chrome/51.0.2704.103 Safari/537.36'}
	request = urllib2.Request(url, None, headers)
	page 	= urllib2.urlopen(request)
	html 	= page.read()

	url_zap = re.findall(r'<a class="catalog-product catalog-item" href="/([^<]+)"', html)
	prefix  = 'http://simple.ripley.cl/'

	# Se revisa el URL de cada zapatilla, debido a que en el URL principal no se encuentran las tallas
	for i in range(0,len(url_zap)):
		# Abrir el URL de cada zapatilla para obtener sus caracteristicas 
		page_zap = urllib2.urlopen(prefix + url_zap[i])
		html_zap = page_zap.read()

		# Encontrar sus tallas disponibles
		tallas = re.findall(r'data-value="([^ ="a-zA-Z]+)', html_zap)
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
		modelo = re.findall(r'<section class="product-header"><h1 itemprop="name">([^<]+)</h1>', html_zap)

		# Encontrar su precio
		precio = re.findall(r'<span class="product-price" itemprop="lowPrice" content="[0-9]+">([0-9\.\$]+)</span>', html_zap)
		if len(precio) == 0: # Algunos productos no tienen precio, al no existir stock al momento de consultar
			continue 

		# Almacenamiento de datos correspondientes
		Modelos_Z.append(modelo)
		Tallas_Z.append(tallas_disp)
		Precios_Z.append(precio)

		# Solo se utiliza en caso de querer ver las URLs de cada zapatilla elegida
		print url_zap[i]
			
	return Modelos_Z, Tallas_Z, Precios_Z	