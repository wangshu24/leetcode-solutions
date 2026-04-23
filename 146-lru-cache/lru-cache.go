type Node struct {
    key, val int
    left, right *Node
}

type LRUCache struct {
    nodes map[int]*Node
    head, tail *Node
    capt int
    leng int
}


func Constructor(capacity int) LRUCache {
    head:= &Node{0,0,nil,nil}
    tail:= &Node{0,0,nil,nil}
    head.left = tail
    tail.right = head
    return LRUCache {
        nodes: make(map[int]*Node),
        head: head,
        tail: tail,
        capt: capacity,
        leng: 0,
    }
}

func remove(n *Node) {
    n.left.right = n.right
    n.right.left = n.left
}

func (this *LRUCache) moveToHead(n *Node) {
    n.right = this.head
    n.left = this.head.left
    this.head.left.right = n
    this.head.left = n

}


func (this *LRUCache) Get(key int) int {
    if node, there := this.nodes[key]; there {
        remove(node)
        this.moveToHead(node)
        return node.val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node, there := this.nodes[key]; there {
        remove(node)
        this.moveToHead(node)
        node.val = value
    } else {
        if len(this.nodes) >= this.capt {
            lrunode := this.tail.right
            remove(lrunode)
            delete(this.nodes, lrunode.key)
        }
        newNode := &Node{key: key, val: value}
        this.nodes[key] = newNode
        this.moveToHead(newNode)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */