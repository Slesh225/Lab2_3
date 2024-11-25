package main

import (
	"bufio"
	"fmt"
	"os"
)

// Array представляет структуру массива
type Array struct {
	MaxCapacity int
	Size        int
	Data        []string
}

// NewArray создает новый массив
func NewArray(capacity int) *Array {
	return &Array{
		MaxCapacity: capacity,
		Size:        0,
		Data:        make([]string, capacity),
	}
}

// Add вставляет элемент по указанному индексу
func (a *Array) Add(index int, value string) {
	if index < 0 || index > a.Size || a.Size >= a.MaxCapacity {
		fmt.Println("Неверный индекс или массив заполнен")
		return
	}
	// Сдвигаем элементы вправо, начиная с указанного индекса
	for i := a.Size; i > index; i-- {
		a.Data[i] = a.Data[i-1]
	}
	a.Data[index] = value // Вставляем элемент
	a.Size++
}

// AddToTheEnd добавляет элемент в конец массива
func (a *Array) AddToTheEnd(value string) {
	if a.Size >= a.MaxCapacity {
		fmt.Println("Массив заполнен")
		return
	}
	a.Data[a.Size] = value // Вставляем элемент в конец
	a.Size++
}

// Remove удаляет элемент по указанному индексу
func (a *Array) Remove(index int) {
	if index < 0 || index >= a.Size {
		fmt.Println("Неверный индекс")
		return
	}
	// Сдвигаем элементы влево, начиная с указанного индекса
	for i := index; i < a.Size-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	a.Size--
}

// Replace заменяет элемент по указанному индексу
func (a *Array) Replace(index int, value string) {
	if index < 0 || index >= a.Size {
		fmt.Println("Неверный индекс")
		return
	}
	a.Data[index] = value // Заменяем элемент
}

// Print выводит элементы массива
func (a *Array) Print() {
	for i := 0; i < a.Size; i++ {
		fmt.Print(a.Data[i], " ")
	}
	fmt.Println()
}

// Length возвращает количество элементов в массиве
func (a *Array) Length() int {
	return a.Size
}

// SaveToFile сохраняет массив в файл
func (a *Array) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	for i := 0; i < a.Size; i++ {
		_, err := file.WriteString(a.Data[i] + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
	}
	return nil
}

// LoadFromFile загружает массив из файла
func (a *Array) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	a.Size = 0 // Очищаем массив перед загрузкой
	for scanner.Scan() && a.Size < a.MaxCapacity {
		a.Data[a.Size] = scanner.Text()
		a.Size++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// Get возвращает элемент по указанному индексу
func (a *Array) Get(index int) string {
	if index < 0 || index >= a.Size {
		fmt.Println("Неверный индекс")
		return ""
	}
	return a.Data[index] // Возвращаем элемент по индексу
}
