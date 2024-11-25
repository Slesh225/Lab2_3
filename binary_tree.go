package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// BinaryTree представляет структуру бинарного дерева
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree создает новое бинарное дерево
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Insert добавляет новый узел в бинарное дерево
func (bt *BinaryTree) Insert(digit int) {
	bt.Root = bt.insert(bt.Root, digit)
}

func (bt *BinaryTree) insert(node *TreeNode, digit int) *TreeNode {
	if node == nil {
		return &TreeNode{Digit: digit}
	}

	if digit < node.Digit {
		node.Left = bt.insert(node.Left, digit)
	} else if digit > node.Digit {
		node.Right = bt.insert(node.Right, digit)
	}

	return node
}

// FindValue ищет значение в бинарном дереве
func (bt *BinaryTree) FindValue(value int) bool {
	return bt.findValue(bt.Root, value)
}

func (bt *BinaryTree) findValue(current *TreeNode, value int) bool {
	if current == nil {
		return false
	}
	if current.Digit == value {
		return true
	}
	return bt.findValue(current.Left, value) || bt.findValue(current.Right, value)
}

// FindIndex находит значение по конкретному индексу
func (bt *BinaryTree) FindIndex(index int) {
	if index < 0 {
		fmt.Println("Неверный индекс.")
		return
	}

	if bt.Root == nil {
		fmt.Println("Дерево пустое.")
		return
	}

	queue := NewQueueTree()
	queue.Enqueue(bt.Root)
	currentIndex := 0

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		if currentIndex == index {
			fmt.Println("Значение:", current.Digit)
			return
		}
		currentIndex++

		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	fmt.Println("Значение не найдено.")
}

// Display печатает бинарное дерево
func (bt *BinaryTree) Display() {
	if bt.Root == nil {
		fmt.Println("Дерево пустое.")
		return
	}
	bt.printCBT(bt.Root, 0)
}

func (bt *BinaryTree) printCBT(current *TreeNode, level int) {
	if current != nil {
		bt.printCBT(current.Right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Println(current.Digit)
		bt.printCBT(current.Left, level+1)
	}
}

// Clear удаляет все узлы из бинарного дерева
func (bt *BinaryTree) Clear() {
	bt.clear(bt.Root)
	bt.Root = nil
}

func (bt *BinaryTree) clear(node *TreeNode) {
	if node != nil {
		bt.clear(node.Left)
		bt.clear(node.Right)
		node.Left = nil
		node.Right = nil
	}
}

// LoadFromFile загружает бинарное дерево из файла
func (bt *BinaryTree) LoadFromFile(file string) error {
	bt.Clear()
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return fmt.Errorf("недопустимое значение в файле: %v", err)
		}
		bt.Insert(value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// SaveToFile сохраняет бинарное дерево в файл
func (bt *BinaryTree) SaveToFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer f.Close()

	queue := NewQueueTree()
	queue.Enqueue(bt.Root)
	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if _, err := f.WriteString(fmt.Sprintf("%d\n", current.Digit)); err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}

		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return nil
}
