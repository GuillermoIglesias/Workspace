import urllib2 
import re 
import sys

# PARIS #
url  = 'http://www.paris.cl/webapp/wcs/stores/servlet/SearchDisplay?searchTermScope=&searchType=1000&filterTerm=&orderBy=&maxPrice=&showResultsPage=true&langId=-5&beginIndex=0&sType=SimpleSearch&metaData=YWRzX2YyNTUwMV9udGtfY3M6IlJ1bm5pbmci&manufacturer=&resultCatEntryType=&catalogId=40000000629&pageView=image&searchTerm=&minPrice=&categoryId=51386230&storeId=10801&pageSize=400'
page = urllib2.urlopen(url)
html = page.read()

url_zap = re.findall(r'<a  href="http://www.paris.cl/tienda/es/paris/moda-calzado/moda-calzado/.*-',html)

for i in range(0,len(url_zap)):
	print str(i) + " " + url_zap[i][10:]

for i in range(0,len(url_zap)):
	page_zap = urllib2.urlopen(url_zap[i][10:])
	html_zap = page_zap.read()

	modelo      = re.findall(r'<meta property="og:description" content="([^<]+)"', html_zap)
	tallas      = re.findall(r'Talla_CL ([^ _|:"a-zA-Z]+)', html_zap)
	if len(tallas) == 0:
		continue
	precio 		= re.findall(r'<div id="price" class="price offerPrice bold">(.+?)</div>', html_zap, flags = re.DOTALL)
	if len(precio) == 0:
		precio 		= re.findall(r'<div id="offerPrice" class="price offerPrice bold">(.+?)</div>', html_zap, flags = re.DOTALL)
	precio_str  = re.findall(r'([0-9\.\$]+)', precio[0])

	print "############## "+str(i) + " ##############"
	print modelo
	print tallas
	print precio_str


