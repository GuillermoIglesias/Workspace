# Tarea 2 CIT2001 Universidad Diego Portales
# Autor: Guillermo Iglesias Birkner
# Profesor: Diego Caro - Ayudante: Farid Abulias

weight = [1,4,3,11,3]
price  = [4,5,1,4,4]

K = len(price)
W = 10

# Solucion al problema con programacion dinamica
# Acercamiento Bottom-Up en O(n^2) o mejor dicho O(K*W)
# Esta funcion se utiliza para construir la tabla de soluciones
def Max_Cargo_Table(weight, price, W, Table):
    for i in range(1,len(weight)+1):
      for j in range(W+1):
        if j >= weight[i-1]:
          Table[i][j] = max(Table[i-1][j],Table[i-1][j-weight[i-1]]+price[i-1])
        else:
          Table[i][j] = Table[i-1][j]

# Acceso a la tabla para encontrar la solucion al problema
def Solution(weight, K, W, Table):
  cap = W
  objects = []
  for i in range(K,0,-1):
    if Table[i][cap] == Table[i-1][cap]:
      continue
    else:
      objects.append(weight[i-1])
      cap = cap - weight[i-1]
  return objects

# Tabla donde se almacenaran las soluciones
Table = []
for i in range(K+1):
  row = []
  for j in range(W+1):
    row.append(0)
  Table.append(row)

Max_Cargo_Table(weight,price,W,Table)

# Imprime Tabla
for i in range(K+1):
  print Table[i]

print Solution(weight,K,W,Table)