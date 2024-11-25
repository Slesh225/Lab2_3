package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filename, query string

	// Проверяем аргументы командной строки
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
			i++
		}
		if arg == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
			i++
		}
	}

	// Если указаны аргументы --file и --query, работаем в режиме командной строки
	if filename != "" && query != "" {
		set := NewSet()

		if err := set.LoadFromFile(filename); err != nil {
			fmt.Printf("Ошибка загрузки множества из файла: %v\n", err)
			return
		}

		set.ProcessQuery(query)

		if err := set.SaveToFile(filename); err != nil {
			fmt.Printf("Ошибка сохранения множества в файл: %v\n", err)
			return
		}
		return
	}

	// Иначе работаем в интерактивном режиме
	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1 - Инфиксная запись")
		fmt.Println("3 - Работа с множеством")
		fmt.Println("4 - Работа с массивами")
		fmt.Println("5 - Работа с бинарными деревьями")
		fmt.Println("6 - Работа с хеш-таблицами")
		fmt.Println("0 - Выход из программы")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue // Пропускаем пустой ввод
		}

		switch input {
		case "1":
			infixToPostfix()
		case "3":
			setOperations()
		case "4":
			arrayOperations()
		case "5":
			binaryTreeOperations()
		case "6":
			hashTableOperations()
		case "0":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите 1, 3, 4, 5, 6 или 0.")
		}
	}
}

func infixToPostfix() {
	fmt.Println("Введите инфиксное выражение:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	postfix := convertInfixToPostfix(input)
	fmt.Println("Постфиксная запись:", postfix)
}

func convertInfixToPostfix(infix string) string {
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	var stack Stack
	var postfix strings.Builder

	for _, char := range infix {
		switch {
		case char >= 'A' && char <= 'Z':
			postfix.WriteRune(char)
		case char == '(':
			stack.Push(string(char))
		case char == ')':
			for stack.Top != nil && stack.Top.Data != "(" {
				postfix.WriteString(stack.Top.Data)
				stack.Pop()
			}
			stack.Pop() // Удаляем '(' из стека
		default:
			for stack.Top != nil && stack.Top.Data != "(" && precedence[rune(stack.Top.Data[0])] >= precedence[char] {
				postfix.WriteString(stack.Top.Data)
				stack.Pop()
			}
			stack.Push(string(char))
		}
	}

	for stack.Top != nil {
		postfix.WriteString(stack.Top.Data)
		stack.Pop()
	}

	return postfix.String()
}

func arrayOperations() {
	fmt.Println("Введите элементы массива через пробел:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	elements := strings.Split(input, " ")
	array := NewArray(len(elements))

	for _, elem := range elements {
		array.AddToTheEnd(elem)
	}

	fmt.Println("Все различные подмассивы:")
	printSubarrays(array)
}

func printSubarrays(array *Array) {
	n := array.Length()
	for i := 0; i < (1 << n); i++ {
		fmt.Print("{")
		for j := 0; j < n; j++ {
			if (i & (1 << j)) > 0 {
				fmt.Print(array.Get(j))
			}
		}
		fmt.Println("}")
	}
}

func hashTableOperations() {
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	maxLength, maxSubstring := lengthOfLongestSubstring(input)
	fmt.Printf("Максимальная длина подстроки с уникальными символами: %d (%s)\n", maxLength, maxSubstring)
}

func lengthOfLongestSubstring(s string) (int, string) {
	charIndexMap := NewHashTable(len(s))
	maxLength := 0
	start := 0
	maxStart := 0

	for i, char := range s {
		charStr := string(char)
		if index, exists := charIndexMap.HGet(charStr); exists {
			indexInt := index.(int)
			if indexInt >= start {
				start = indexInt + 1
			}
		}
		charIndexMap.HSet(charStr, i)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			maxStart = start
		}
	}

	return maxLength, s[maxStart : maxStart+maxLength]
}

func binaryTreeOperations() {
	binaryTree := NewBinaryTree()

	for {
		fmt.Println("Введите число (для завершения введите 'q'):")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		digit, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Неверный ввод. Пожалуйста, введите число.")
			continue
		}

		if !binaryTree.FindValue(digit) {
			binaryTree.Insert(digit)
		}
	}

	height := binaryTree.Height()
	fmt.Printf("Высота дерева: %d\n", height)
}

// Height возвращает высоту бинарного дерева
func (bt *BinaryTree) Height() int {
	return bt.height(bt.Root)
}

func (bt *BinaryTree) height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := bt.height(node.Left)
	rightHeight := bt.height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
