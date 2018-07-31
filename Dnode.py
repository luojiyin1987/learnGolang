class Dnode(object):
    def __init__(self, value=0, prev=None, next=None):
        self.value = value
        self.next = next
        self.prev = prev


class Dlink(object):
    def __init__(self):
        self.head = None
    
    def search(self, key):
        x = self.head
        while x != None and x.value != key:
            x = x.next
        return x

    def insert(self, x):
        x.next = x.head
        if self.head is not None:
            self.head.prev = x
        self.head = x 
        x.prev = None

    def delete(self, x):
        if not isinstance(x, Dnode): x =self.search(x)
        
        if x is None: return 

        if x.prev != None: 
            x.prev.next = x.next
        
        else:
            self.head = x.next
        
        if x.next != None:
            x.next.prev = x.prev

    def show(self):
        x = self.head
        while x != None:
            print(x.value)
            x = x.next
