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
mem = {}
def ks(index, weight):
    global items
    global c
    global mem
    c += 1
 
    case = (index, weight)
    if case in mem:
        return mem[case]
 
 
    if index >= len(items):
        mem[case] = 0
        return mem[case]
 
    item = items[index]
    grab_case = (index + 1, weight - item.weight)
    no_grab_case = (index + 1, weight)
 
    if item.weight > weight:
        mem[case] = ks(*no_grab_case)
        return mem[case]
    else:
        if no_grab_case not in mem:
            mem[no_grab_case] = ks(index + 1, weight)
        if grab_case not in mem:
            mem[grab_case] = ks(index + 1, weight - item.weight) + item.value
        mem[case] = max(mem[grab_case], mem[no_grab_case])
        return mem[case]
 
print "Max sum: %d" % (ks(0, 20),)
print "Iterations %d" % (c,)
print