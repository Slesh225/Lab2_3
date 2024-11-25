package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HashNode представляет узел в хэш-таблице
type HashNode struct {
	Key   string
	Value interface{}
	Next  *HashNode
}

// HashTable представляет структуру хэш-таблицы
type HashTable struct {
	Capacity int
	Table    []*HashNode
	size     int // Добавляем поле для хранения размера
}

// NewHashTable создает новую хэш-таблицу
func NewHashTable(size int) *HashTable {
	return &HashTable{
		Capacity: size,
		Table:    make([]*HashNode, size),
		size:     0, // Инициализируем размер
	}
}

// HashFunction вычисляет индекс хэша для заданного ключа
func (ht *HashTable) HashFunction(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (hash*31 + int(ch)) % ht.Capacity
	}
	return hash
}

// HSet вставляет или обновляет пару ключ-значение в хэш-таблице
func (ht *HashTable) HSet(key string, value interface{}) {
	index := ht.HashFunction(key)
	current := ht.Table[index]

	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		current = current.Next
	}

	newNode := &HashNode{Key: key, Value: value, Next: ht.Table[index]}
	ht.Table[index] = newNode
	ht.size++ // Увеличиваем размер при добавлении нового элемента
}

// HGet извлекает значение, связанное с ключом
func (ht *HashTable) HGet(key string) (interface{}, bool) {
	index := ht.HashFunction(key)
	current := ht.Table[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return nil, false
}

// HDel удаляет пару ключ-значение из хэш-таблицы
func (ht *HashTable) HDel(key string) {
	index := ht.HashFunction(key)
	current := ht.Table[index]
	var prev *HashNode

	for current != nil {
		if current.Key == key {
			if prev == nil {
				ht.Table[index] = current.Next
			} else {
				prev.Next = current.Next
			}
			ht.size-- // Уменьшаем размер при удалении элемента
			return
		}
		prev = current
		current = current.Next
	}
}

// Clear удаляет все элементы из хэш-таблицы
func (ht *HashTable) Clear() {
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			temp := current
			current = current.Next
			temp.Next = nil
		}
		ht.Table[i] = nil
	}
	ht.size = 0 // Сбрасываем размер
}

// HPrint печатает содержимое хэш-таблицы
func (ht *HashTable) HPrint() {
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		if current != nil {
			fmt.Printf("[%d]: ", i)
			for current != nil {
				fmt.Printf("%s => %v ", current.Key, current.Value)
				current = current.Next
			}
			fmt.Println()
		}
	}
}

// HExists проверяет, существует ли ключ в хэш-таблице
func (ht *HashTable) HExists(key string) bool {
	index := ht.HashFunction(key)
	current := ht.Table[index]

	for current != nil {
		if current.Key == key {
			return true
		}
		current = current.Next
	}

	return false
}

// LoadFromFile загружает хэш-таблицу из файла
func (ht *HashTable) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			value, err := strconv.ParseBool(parts[1])
			if err != nil {
				return fmt.Errorf("ошибка преобразования значения: %v", err)
			}
			ht.HSet(parts[0], value)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// SaveToFile сохраняет хэш-таблицу в файл
func (ht *HashTable) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			_, err := file.WriteString(fmt.Sprintf("%s %t\n", current.Key, current.Value))
			if err != nil {
				return fmt.Errorf("ошибка записи в файл: %v", err)
			}
			current = current.Next
		}
	}
	return nil
}

// Size возвращает количество элементов в хэш-таблице
func (ht *HashTable) Size() int {
	return ht.size
}

// ForEach выполняет функцию для каждого элемента в хэш-таблице
func (ht *HashTable) ForEach(fn func(key string, value interface{})) {
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			fn(current.Key, current.Value)
			current = current.Next
		}
	}
}
