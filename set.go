package main

import (
	"fmt"
	"strings"
)

// Set представляет структуру множества
type Set struct {
	elements *HashTable
}

// NewSet создает новое множество
func NewSet() *Set {
	return &Set{
		elements: NewHashTable(16), // Инициализация с начальной емкостью 16
	}
}

// Add добавляет элемент в множество
func (s *Set) Add(element string) {
	s.elements.HSet(element, true)
}

// Delete удаляет элемент из множества
func (s *Set) Delete(element string) {
	s.elements.HDel(element)
}

// Contains проверяет, является ли элемент частью множества
func (s *Set) Contains(element string) bool {
	_, exists := s.elements.HGet(element)
	return exists
}

// LoadFromFile загружает множество из файла
func (s *Set) LoadFromFile(filename string) error {
	return s.elements.LoadFromFile(filename)
}

// SaveToFile сохраняет множество в файл
func (s *Set) SaveToFile(filename string) error {
	return s.elements.SaveToFile(filename)
}

// ProcessQuery обрабатывает команды для множества
func (s *Set) ProcessQuery(query string) {
	tokens := strings.Split(query, " ")

	switch tokens[0] {
	case "SETADD":
		if len(tokens) == 2 {
			s.Add(tokens[1])
			fmt.Printf("Элемент %s добавлен в множество.\n", tokens[1])
		} else {
			fmt.Println("Ошибка: команда SETADD требует 1 аргумент.")
		}
	case "SETDEL":
		if len(tokens) == 2 {
			s.Delete(tokens[1])
			fmt.Printf("Элемент %s удален из множества.\n", tokens[1])
		} else {
			fmt.Println("Ошибка: команда SETDEL требует 1 аргумент.")
		}
	case "SET_AT":
		if len(tokens) == 2 {
			if s.Contains(tokens[1]) {
				fmt.Printf("Элемент %s является частью множества.\n", tokens[1])
			} else {
				fmt.Printf("Элемент %s не является частью множества.\n", tokens[1])
			}
		} else {
			fmt.Println("Ошибка: команда SET_AT требует 1 аргумент.")
		}
	default:
		fmt.Printf("Неизвестная команда: %s\n", tokens[0])
	}
}
