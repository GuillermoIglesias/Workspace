# Tarea 1 CIT2001 

import Sol_Sencilla
import Sol_Optima
import RNG
import timeit

r = 2 # Modifique este valor para la cantidad de repeticiones del experimento

# Se utilizaron las siguientes muestras
n = [100000,200000,400000,800000,1600000,3200000,6400000] 	

# Retorna los tiempos promedios de cada operacion
def Tiempo(timer1, timer2, r):
	t_sencilla_sum  = timer1.repeat(r,1)
	t_optima_sum 	= timer2.repeat(r,1)

	prom_sencilla 	= sum(t_sencilla_sum)/r
	prom_optima 	= sum(t_optima_sum)/r

	return prom_sencilla, prom_optima

for j in range(len(n)):
	
	k = [1,n[j]/2,n[j]] 		# Valores de kth utilizados para el experimento
	m = [n[j]/10,n[j]/2,n[j]] 	# Valores de Top M utilizados para el experimento

	# Se generan las listas de numeros, los cuales se almacenan en archivos .txt
	RNG.CreateRandom(n[j])
	Random = RNG.ReadRandom(n[j])
	Orden  = RNG.CreateOrden(n[j])

	print "K-esimo termino, soluciones para N =",n[j]
	
	for i in range(len(k)):
		# Mediciones para Kth elemento en una lista Ordenada
		t_sencilla_kth_orden	= timeit.Timer(lambda: Sol_Sencilla.kth_sencillo(Orden,k[i]))
		t_optimo_kth_orden  	= timeit.Timer(lambda: Sol_Optima.kth_optimo(Orden,k[i]))

		# Mediciones para Kth elemento en una lista Aleatoria
		t_sencilla_kth_rnd 		= timeit.Timer(lambda: Sol_Sencilla.kth_sencillo(Random,k[i]))
		t_optimo_kth_rnd   		= timeit.Timer(lambda: Sol_Optima.kth_optimo(Random,k[i]))

		# Almacenamiento de resultados (mediciones de tiempo)
		t_kth_simple_orden, t_kth_op_orden  = Tiempo(t_sencilla_kth_orden, t_optimo_kth_orden, r)
		t_kth_simple_rnd, t_kth_op_rnd 		= Tiempo(t_sencilla_kth_rnd, t_optimo_kth_rnd, r)

		print "Arreglo Ordenado, K =",k[i]
		print "- Solucion Sencilla: ", t_kth_simple_orden
		print "- Solucion Optima:   ", t_kth_op_orden

		print "Arreglo Aleatorio, K =",k[i]
		print "- Solucion Sencilla: ", t_kth_simple_rnd
		print "- Solucion Optima:   ", t_kth_op_rnd

	print "Top-M terminos, soluciones para N =",n[j]
	
	for i in range(len(m)):
		# Mediciones para los Top M elementos en una lista Ordenada
		t_sencilla_topM_orden	= timeit.Timer(lambda: Sol_Sencilla.top_m_sencillo(Orden,m[i]))
		t_optimo_topM_orden  	= timeit.Timer(lambda: Sol_Optima.top_m_optimo(Orden,m[i]))

		# Mediciones para los Top M elementos en una lista Aleatoria
		t_sencilla_topM_rnd 		= timeit.Timer(lambda: Sol_Sencilla.top_m_sencillo(Random,m[i]))
		t_optimo_topM_rnd   		= timeit.Timer(lambda: Sol_Optima.top_m_optimo(Random,m[i]))

		# Almacenamiento de resultados (mediciones de tiempo)
		t_topM_simple_orden, t_topM_op_orden    = Tiempo(t_sencilla_topM_orden, t_optimo_topM_orden, r)
		t_topM_simple_rnd, t_topM_op_rnd 		= Tiempo(t_sencilla_topM_rnd, t_optimo_topM_rnd, r)

		print "Arreglo Ordenado, Top",m[i],"elementos."
		print "- Solucion Sencilla: ", t_topM_simple_orden
		print "- Solucion Optima:   ", t_topM_op_orden

		print "Arreglo Aleatorio, Top",m[i],"elementos."
		print "- Solucion Sencilla: ", t_topM_simple_rnd
		print "- Solucion Optima:   ", t_topM_op_rnd
