package main

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

// 注意，头、尾节点是不会变的，他们只是边界标记
type LRUCache struct {
	Cap   int
	Cache map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	Cache := make(map[int]*Node)
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		Cap:   capacity,
		Cache: Cache,
		Head:  head,
		Tail:  tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.Cache[key]; ok {
		// node移到头部
		l.MoveToHead(node)
		return node.Value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.Cache[key]; ok {
		l.Cache[key].Value = value
		l.MoveToHead(node)
		// 如果不用else，这里要return
	} else {
		// 添加新节点到头部
		newNode := &Node{
			Key:   key,
			Value: value,
		}
		l.AddToHead(newNode)
		if len(l.Cache) > l.Cap {
			// 删除尾巴
			l.RemoveTail()
		}
	}
}

func (l *LRUCache) MoveToHead(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Pre = l.Head
	node.Next = l.Head.Next
	l.Head.Next.Pre = node // 关键步骤
	l.Head.Next = node
}

func (l *LRUCache) AddToHead(node *Node) {
	l.Cache[node.Key] = node
	node.Next = l.Head.Next
	node.Pre = l.Head
	l.Head.Next.Pre = node
	l.Head.Next = node
}

func (l *LRUCache) RemoveTail() {
	// tail的前一个节点才是真正的尾
	tail := l.Tail.Pre
	tail.Pre.Next = l.Tail
	l.Tail.Pre = tail.Pre
	tail.Pre = nil
	tail.Next = nil
	delete(l.Cache, tail.Key)
}
