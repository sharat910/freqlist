package freqlist

import "fmt"

type FreqList struct {
	head         *Node
	tail         *Node
	maxSize      int
	curSize      int
	nodesCreated int
	nodesDeleted int
}

func (l *FreqList) PrintStats() {
	fmt.Printf("Nodes created: %d\n", l.nodesCreated)
	fmt.Printf("Nodes deleted: %d\n", l.nodesDeleted)
	fmt.Printf("Current size: %d\n", l.curSize)
}
func New(maxSize int) *FreqList {
	return &FreqList{maxSize: maxSize}
}

type Node struct {
	key  interface{}
	freq int
	prev *Node
	next *Node
}

func (n *Node) Freq() int {
	return n.freq
}

func (l *FreqList) NewNode(key interface{}) (newNode *Node, deletedNodeKey interface{}) {
	newNode = &Node{key: key}
	l.curSize++
	l.nodesCreated++
	if l.curSize > l.maxSize {
		deletedNodeKey = l.tail.key
		l.RemoveNode(l.tail)
		l.nodesDeleted++
	}
	// Empty list
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	// Head is still not accessed
	if l.head.freq == 0 {
		l.head.next = newNode
		newNode.prev = l.head
		l.head = newNode
		return
	}

	if l.tail.freq != 0 {
		l.tail.prev = newNode
		newNode.next = l.tail
		l.tail = newNode
		return
	}

	// Find the node to insert before
	// Among the nodes with the same frequency, insert at the top -- new nodes are accessed more recently
	insertBeforeThisNode := l.tail
	for insertBeforeThisNode.freq == 0 {
		insertBeforeThisNode = insertBeforeThisNode.next
	}
	newNode.prev = insertBeforeThisNode.prev
	insertBeforeThisNode.prev.next = newNode
	insertBeforeThisNode.prev = newNode
	newNode.next = insertBeforeThisNode
	return
}

func (l *FreqList) AccessNode(n *Node) {
	n.freq++
	if l.head == n {
		return
	}

	// Move node to the right position
	for n.freq > n.next.freq {
		curPrev := n.prev
		curNext := n.next
		n.prev = curNext
		n.next = curNext.next
		curNext.next = n
		curNext.prev = curPrev
		if curPrev != nil {
			curPrev.next = curNext
		} else {
			l.tail = curNext
		}
		if n.next != nil {
			n.next.prev = n
		} else {
			l.head = n
			break
		}
	}
}

func (l *FreqList) RemoveNode(n *Node) {
	l.curSize--
	if n == l.head {
		l.head = n.prev
	}
	if n == l.tail {
		l.tail = n.next
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}
