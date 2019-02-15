# Tarea 2 CIT2001

import RNG
import Solucion
import timeit

W         = 500     # Peso total del camion para transportar
r         = 1       # Modifique este valor para la cantidad de repeticiones del experimento
PrintSol  = True    # Modifique este valor si se quiere imprimir la solucion [True,False]
KMax      = 10250   # Modifique este valor para el K Maximo a considerar

# Se realizan 20 experimentos, con distintos valores de K
for K in range(250,KMax,250):
  # Se generan las listas de numeros, los cuales se almacenan en archivos .txt
  RNG.CreateRandom(K)
  name, weight, price = RNG.ReadRandom(K)

  # Timer para contabilizar el tiempo de ejecucion de la solucion
  t_timer = timeit.Timer(lambda: Solucion.Max_Cargo_Table(name,weight,price,K,W,PrintSol))
  t_sum   = t_timer.repeat(r,1)
  avg_t   = sum(t_sum)/r

  # Imprime tiempo prom de solucion
  print "# Tiempo Promedio (K = " + str(K) + "): " + str(avg_t)