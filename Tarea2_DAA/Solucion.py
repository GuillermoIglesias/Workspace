# Solucion al problema con programacion dinamica
# Acercamiento Bottom-Up en O(K*W)

# Esta funcion se utiliza para construir la tabla de soluciones
def Max_Cargo_Table(name, weight, price, K, W, PrintSol):
  # Tabla donde se almacenaran las soluciones
  Table = []
  for i in range(K+1):
    row = []
    for j in range(W+1):
      row.append(0)
    Table.append(row)

  # Relleno de la tabla de soluciones  
  for i in range(1,len(weight)+1):
    for j in range(W+1):
      if j >= weight[i-1]:
        Table[i][j] = max(Table[i-1][j],Table[i-1][j-weight[i-1]]+price[i-1])
      else:
        Table[i][j] = Table[i-1][j]  
      
  Solution = Find_Solution(weight,K,W,Table)

  # Se utiliza en caso de querer imprimir la solucion
  # Su tiempo de ejecucion es O(len(Solution))
  # Pero dado su valor tan bajo, es despreciable con respecto al K y W entregado
  if PrintSol == True:
    Print_Solution(name,weight,price,K,W,Solution,Table)


# Acceso a la tabla para encontrar la solucion al problema
def Find_Solution(weight, K, W, Table):
  cap = W
  objects = []
  for i in range(K,0,-1):
    if Table[i][cap] == Table[i-1][cap]:
      continue
    else:
      objects.append(i-1)
      cap = cap - weight[i-1]
  return objects

# Item # | peso | precio
def Print_Solution(name, weight, price, K, W, Solution, Table):
  for i in range(len(Solution)):
    print name[Solution[i]] + " | peso: " + str(weight[Solution[i]]) + "Kg | precio: $" + str(price[Solution[i]])  
  print "valor total: $" + str(Table[K][W])
