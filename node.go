package main

// Node представляет узел в односвязном списке
type Node struct {
	Data string
	Next *Node
}

// DoublyNode представляет узел в двусвязном списке
type DoublyNode struct {
	Data string
	Next *DoublyNode
	Prev *DoublyNode
}

// TreeNode представляет узел в бинарном дереве
type TreeNode struct {
	Digit int
	Left  *TreeNode
	Right *TreeNode
}

// QueueNode представляет узел в очереди для узлов дерева
type QueueNode struct {
	Tree *TreeNode
	Next *QueueNode
}

// QueueTree представляет очередь для узлов дерева
type QueueTree struct {
	Front *QueueNode
	Rear  *QueueNode
	Count int
}

// NewQueueTree создает новую очередь для узлов дерева
func NewQueueTree() *QueueTree {
	return &QueueTree{}
}

// IsEmpty проверяет, пуста ли очередь
func (qt *QueueTree) IsEmpty() bool {
	return qt.Count == 0
}

// Enqueue добавляет узел дерева в очередь
func (qt *QueueTree) Enqueue(node *TreeNode) {
	newNode := &QueueNode{Tree: node}
	if qt.Rear == nil {
		qt.Front = newNode
		qt.Rear = newNode
	} else {
		qt.Rear.Next = newNode
		qt.Rear = newNode
	}
	qt.Count++
}

// Dequeue удаляет узел дерева из очереди
func (qt *QueueTree) Dequeue() *TreeNode {
	if qt.IsEmpty() {
		return nil
	}

	newNode := qt.Front
	res := qt.Front.Tree
	qt.Front = qt.Front.Next

	if qt.Front == nil {
		qt.Rear = nil
	}

	newNode.Next = nil
	qt.Count--
	return res
}
