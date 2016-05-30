def mochila(w,p,capacidadW,T):
    #for c in range(capacidadW+1):
     # T[0][c]=0
    for j in range(1,len(w)+1):
      for c in range(capacidadW+1):
        if c >= w[j-1]:
            T[j][c]=max(T[j-1][c],T[j-1][c-w[j-1]]+p[j-1])
        else:
            T[j][c]=T[j-1][c]
    for i in range(len(w)+1):
        print T[i]

def printPath(n, capacidadW, w):
  j = capacidadW
  sol = []
  for i in range(n,0,-1):
    if T[i][j] == T[i-1][j]:
      continue
    else:
      sol.append(w[i-1])
      j = j - w[i-1]
  return sol

def Max_Cargo_Table(weight, price, W, Table):
    for i in range(1,len(weight)+1):
      for j in range(W+1):
        if j >= weight[i-1]:
          Table[i][j] = max(Table[i-1][j],Table[i-1][j-weight[i-1]]+price[i-1])
        else:
          Table[i][j] = Table[i-1][j]
                

w=[1,4,3,11,3]
p=[4,5,1,4,4]
capacidadW=10

T = []
for i in range(len(w)+1):
  ea_row = []
  for j in range(capacidadW+1):
    ea_row.append(0)
  T.append(ea_row)

for i in range(len(w)+1):
  print T[i]


mochila(w,p,capacidadW,T) 
print printPath(len(p),capacidadW,w)

Table = [[0]*(capacidadW+1)]*(len(w)+1)

Max_Cargo_Table(w,p,capacidadW,T)

for i in range(len(w)+1):
  print T[i]

