class Item:
    def __init__(self, weight, value):
        self.weight = weight
        self.value = value
         
    def __str__(self):
        return "Peso: %d - Valor: %d" % (self.weight, self.value)
    
    def __repr__(self):
        return self.__str__()
 
items = [Item(1, 4), Item(4, 5), Item(3, 1), Item(11, 4), Item(3, 4),]

c = 0
def ks(index, weight):
    global items
    global c
    c += 1
 
    if index >= len(items):
        return 0
 
    item = items[index]
 
    if item.weight > weight:
        return ks(index + 1, weight)
    else:
        return max(ks(index + 1, weight),
                   ks(index + 1, weight - item.weight) + item.value)
 
print "Max sum: %d" % (ks(0, 20),)
print "Iterations %d" % (c,)
print