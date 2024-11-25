package main

import (
	"fmt"
)

const MAX_SIZE = 100 // Максимальный размер множества

func printSubsets(nums []int, target int, numSubsets int) {
	sets := make([]*Set, numSubsets) // Создаем массив множеств для хранения подмножеств
	for i := 0; i < numSubsets; i++ {
		sets[i] = NewSet()
	}

	var backtrack func(index int, currentSums []int) bool
	backtrack = func(index int, currentSums []int) bool {
		if index == len(nums) {
			// Проверяем, что все подмножества имеют нужную сумму
			for i := 0; i < numSubsets; i++ {
				if currentSums[i] != target {
					return false
				}
			}
			// Выводим подмножества
			fmt.Printf("Разбиение на %d подмножеств с суммой %d:\n", numSubsets, target)
			for i := 0; i < numSubsets; i++ {
				fmt.Print("{ ")
				sets[i].elements.ForEach(func(key string, value interface{}) {
					fmt.Printf("%s ", key)
				})
				fmt.Printf("} Сумма: %d\n", target)
			}
			return true
		}

		// Пытаемся поместить nums[index] в каждое подмножество
		for i := 0; i < numSubsets; i++ {
			if currentSums[i]+nums[index] <= target {
				sets[i].Add(fmt.Sprintf("%d", nums[index]))
				currentSums[i] += nums[index]
				if backtrack(index+1, currentSums) {
					return true
				}
				// Откатываем изменения
				sets[i].Delete(fmt.Sprintf("%d", nums[index]))
				currentSums[i] -= nums[index]
			}
		}
		return false
	}

	// Инициализируем суммы для каждого подмножества
	currentSums := make([]int, numSubsets)
	if backtrack(0, currentSums) {
		return
	} else {
		fmt.Printf("Невозможно разбить множество на %d подмножеств с суммой %d.\n", numSubsets, target)
	}
}

func setOperations() {
	fmt.Printf("Введите количество элементов в множестве (максимум %d): ", MAX_SIZE)
	var size int
	fmt.Scan(&size)
	if size > MAX_SIZE {
		fmt.Println("Количество элементов превышает максимальный размер.")
		return
	}

	fmt.Println("Введите элементы множества (натуральные числа):")
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&elements[i])
	}

	fmt.Println("Введите желаемую сумму для каждого подмножества:")
	var target int
	fmt.Scan(&target)

	totalSum := 0
	for _, num := range elements {
		totalSum += num
	}

	if totalSum%target != 0 {
		fmt.Printf("Невозможно разбить множество на подмножества с суммой %d.\n", target)
		return
	}

	numSubsets := totalSum / target
	printSubsets(elements, target, numSubsets)
}
