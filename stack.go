package main

import (
	"bufio"
	"fmt"
	"os"
)

// Stack представляет структуру стека
type Stack struct {
	Top  *Node
	Size int
}

// NewStack создает новый стек
func NewStack() *Stack {
	return &Stack{}
}

// Push добавляет новый элемент на вершину стека
func (s *Stack) Push(value string) {
	newNode := &Node{Data: value, Next: s.Top}
	s.Top = newNode
	s.Size++
}

// Pop удаляет верхний элемент из стека
func (s *Stack) Pop() {
	if s.Top == nil {
		fmt.Println("Стек пуст.")
		return
	}
	temp := s.Top
	s.Top = s.Top.Next
	temp.Next = nil
	s.Size--
}

// Print выводит элементы стека
func (s *Stack) Print() {
	temp := s.Top
	for temp != nil {
		fmt.Print(temp.Data, " ")
		temp = temp.Next
	}
	fmt.Println()
}

// SaveToFile сохраняет стек в файл
func (s *Stack) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	temp := s.Top
	for temp != nil {
		_, err := file.WriteString(temp.Data + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
		temp = temp.Next
	}
	return nil
}

// LoadFromFile загружает стек из файла
func (s *Stack) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.Push(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}
