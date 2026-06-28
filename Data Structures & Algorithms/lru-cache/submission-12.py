
class Node:
    def __init__(self, key: int, value: int):
        self.next = None
        self.prev = None
        self.value = value
        self.key = key
#    def __str__(self) -> str:
 #       return f"({self.key},{self.value}), prev:{self.prev}, next:{self.next}"

class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.head = None
        self.tail = None
        self.cache = dict()

    def get(self, key: int) -> int:
        if key in self.cache:
            node = self.cache[key]
            if node != self.head:
                self.promote(node)
            return node.value
        return -1
        
    def put(self, key: int, value: int) -> None:
        node = None
        if key in self.cache:
            node = self.cache[key]
            node.value = value
        else:
            node = Node(key, value)
            if len(self.cache)==self.cap:
                self.evictKey()
            self.cache[key]=node
        if node != self.head:
            self.promote(node)
        #print('put of ({},{})', key, value)
        #print('dict: {}', self.cache)
        #if len(self.cache)>self.cap:
            #print('evicting')

    def promote(self, node:Node):
        if self.head:
            # new node
            if node.prev == None and node.next == None:
                node.next = self.head
                self.head.prev = node
                self.head = node
                return 
            if node == self.tail:
                self.tail = node.prev            
                self.tail.next = None
            if node.next: 
                node.next.prev = node.prev
            node.prev.next = node.next
            node.next = self.head
            self.head.prev = node
            self.head = node
        else:
            self.head = node
            self.tail = node

    # def promote(self, node: Node):
    #     #head = self.head
    #     #tail = self.tail
    #     if self.head:
    #         if node.prev:
    #             node.prev.next = node.next
    #         node.next = self.head
    #         node.prev = self.head.prev
    #         self.head.prev = node
    #         if node == self.tail:
    #             self.tail = self.tail.prev
    #     else:
    #         self.tail = node
    #         self.tail.next = node
    #         node.prev = node
    #     self.head = node
    #  #   print('head is {}\ntail is {}', self.head,self.tail)


    def evictKey(self):
        print("evicting")
        if self.tail:
            self.cache.pop(self.tail.key, None)
            self.tail = self.tail.prev
            if self.tail:
                self.tail.next = None
        # if self.tail:
        #     #print('there is tail')
        #     #print('deleted entry from cache')
        #     #print('cache now: {}', self.cache)
        #     self.tail.prev.next = self.tail.next
        #     self.tail.next.prev = self.tail.prev
        #     prev = tail.prev
        #     self.tail.next = None
        #     self.tail.prev = None
        #     self.tail = prev

