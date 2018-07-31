class Queue(object):
    def __init__(self):
        self.queue = []

    def isEmpty(self):
        return self.queue == []

    def EnQueue(self,x):
        self.queue.append(x)
    def DeQueue(self):
        if self.queue:
            a = self.queue[0]
            self.queue.remove(a)
            return a 
        else:
            raise IndexError("It is empty")

    def size(self):
        return len(self.queue)
    

