type Node struct {
    key, val    int
    left, right *Node
}

type LRUCache struct {
    nodes      map[int]*Node
    capacity   int
    head, tail *Node // Sentinel nodes
}

func Constructor(capacity int) LRUCache {
    // Initialize dummy nodes
    head := &Node{0, 0, nil, nil}
    tail := &Node{0, 0, nil, nil}
    head.right = tail
    tail.left = head
    
    return LRUCache{
        nodes:    make(map[int]*Node),
        capacity: capacity,
        head:     head,
        tail:     tail,
    }
}

// Helper to remove a node from the list
func (this *LRUCache) remove(n *Node) {
    n.left.right = n.right
    n.right.left = n.left
}

// Helper to add a node right after the dummy head
func (this *LRUCache) addToHead(n *Node) {
    n.right = this.head.right
    n.left = this.head
    this.head.right.left = n
    this.head.right = n
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.nodes[key]; ok {
        this.remove(node)
        this.addToHead(node)
        return node.val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.nodes[key]; ok {
        this.remove(node)
        node.val = value
        this.addToHead(node)
    } else {
        if len(this.nodes) >= this.capacity {
            // Remove the real tail (the node before dummy tail)
            lruNode := this.tail.left
            this.remove(lruNode)
            delete(this.nodes, lruNode.key)
        }
        newNode := &Node{key: key, val: value}
        this.nodes[key] = newNode
        this.addToHead(newNode)
    }
}