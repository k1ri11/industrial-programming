package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var taskNumber, moduleId int
	fmt.Println("Введите номер модуля")
	fmt.Scanf("%d\n", &moduleId)
	fmt.Println("Введите номер задачи")
	fmt.Scanf("%d\n", &taskNumber)
	xy := fmt.Sprintf("%d_%d", moduleId, taskNumber)
	switch xy {
	case "1_1":
		task1_1()
	case "1_2":
		task1_2()
	case "1_3":
		task1_3()
	case "1_4":
		task1_4()
	case "1_5":
		task1_5()
	case "2_1":
		task2_1()
	case "2_2":
		task2_2()
	case "2_3":
		task2_3()
	case "2_4":
		task2_4()
	case "2_5":
		task2_5()
	case "3_1":
		task3_1()
	case "3_2":
		task3_2()
	case "3_3":
		task3_3()
	case "3_4":
		task3_4()
	case "3_5":
		task3_5()

	default:
		fmt.Println("такая задача не найдена", xy)

	}
}

func task1_1() {
	var number string
	var fromBase, toBase int

	fmt.Println("Введите число, исходную систему счисления и конечную систему счисления:")
	fmt.Scan(&number, &fromBase, &toBase)

	decimal, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		fmt.Println("Ошибка: неверный ввод")
		return
	}

	result := strconv.FormatInt(decimal, toBase)
	fmt.Println("Результат:", result)
}

func task1_2() {
	var a, b, c float64

	fmt.Print("Введите коэффициенты a, b, c через пробел: ")
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Корни уравнения: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("Один корень: x = %.2f\n", x)
	} else {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
		fmt.Printf("Комплексные корни: x1 = %.2f + %.2fi, x2 = %.2f - %.2fi\n", realPart, imaginaryPart, realPart, imaginaryPart)
	}
}

func task1_3() {
	var n int
	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if math.Abs(float64(array[j])) > math.Abs(float64(array[j+1])) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	fmt.Println("Отсортированный массив:", array)
}

func task1_4() {
	var n1, n2 int

	fmt.Print("Введите количество элементов первого массива: ")
	fmt.Scan(&n1)
	arr1 := make([]int, n1)
	fmt.Println("Введите элементы первого массива:")
	for i := 0; i < n1; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Введите количество элементов второго массива: ")
	fmt.Scan(&n2)
	arr2 := make([]int, n2)
	fmt.Println("Введите элементы второго массива:")
	for i := 0; i < n2; i++ {
		fmt.Scan(&arr2[i])
	}

	mergedArray := mergeArrays(arr1, arr2)
	fmt.Println("Слияние массивов:", mergedArray)
}

func mergeArrays(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
}

func task1_5() {
	var str, substr string

	fmt.Print("Введите строку: ")
	fmt.Scan(&str)
	fmt.Print("Введите подстроку: ")
	fmt.Scan(&substr)

	index := findSubstring(str, substr)
	fmt.Printf("Индекс первого вхождения подстроки: %d\n", index)
}

func findSubstring(str, substr string) int {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func task2_1() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите оператор (+, -, *, /, ^, %): ")
	fmt.Scan(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f\n", result)
	}
}

func calculate(a float64, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	case "^":
		return math.Pow(a, b), nil
	case "%":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return math.Mod(a, b), nil
	default:
		return 0, errors.New("недопустимая операция")
	}
}

func task2_2() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	filtered := ""

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered += string(r)
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}
	return true
}

func task2_3() {
	var a1, b1, a2, b2, a3, b3 int

	fmt.Println("Введите начало и конец первого отрезка:")
	fmt.Scan(&a1, &b1)
	fmt.Println("Введите начало и конец второго отрезка:")
	fmt.Scan(&a2, &b2)
	fmt.Println("Введите начало и конец третьего отрезка:")
	fmt.Scan(&a3, &b3)

	if findIntersection(a1, b1, a2, b2, a3, b3) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func findIntersection(a1, b1, a2, b2, a3, b3 int) bool {
	start := max(a1, max(a2, a3))
	end := min(b1, min(b2, b3))

	return start <= end
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func task2_4() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите предложение:")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	fmt.Println("Самое длинное слово:", longestWord(sentence))
}

func longestWord(sentence string) string {
	words := strings.FieldsFunc(sentence, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsDigit(c)
	})

	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func task2_5() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

func task3_1() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	a, b := 0, 1
	for a <= n {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func task3_2() {
	var start, end int
	fmt.Print("Введите два числа (начало и конец диапазона): ")
	fmt.Scan(&start, &end)

	printPrimesInRange(start, end)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func printPrimesInRange(start, end int) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func task3_3() {
	var start, end int
	fmt.Print("Введите два числа (начало и конец диапазона): ")
	fmt.Scan(&start, &end)

	printArmstrongNumbers(start, end)
}

func isArmstrong(num int) bool {
	sum := 0
	temp := num
	digits := int(math.Log10(float64(num))) + 1

	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == num
}

func printArmstrongNumbers(start, end int) {
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func task3_4() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scan(&input)

	reversed := reverseString(input)
	fmt.Println("Строка в обратном порядке:", reversed)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func task3_5() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println("НОД:", a)
}
