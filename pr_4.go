package main

import (
	"fmt"
	"math"
	"strings"
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
		sumDigits()
	case "1_2":
		celsiusToFahrenheit()
	case "1_3":
		doubleArray()
	case "1_4":
		joinStrings()
	case "1_5":
		distance()
	case "2_1":
		isEven()
	case "2_2":
		isLeapYear()
	case "2_3":
		maxOfThree()
	case "2_4":
		ageCategory()
	case "2_5":
		isDivisibleBy3And5()
	case "3_1":
		factorial()
	case "3_2":
		fibonacci()
	case "3_3":
		reverseArray()
	case "3_4":
		primeNumbers()
	case "3_5":
		sumArray()

	default:
		fmt.Println("такая задача не найдена", xy)

	}
}

func sumDigits() {
	var num int
	fmt.Print("Введите число для вычисления суммы его цифр: ")
	fmt.Scan(&num)

	res := 0
	res += num % 10
	num = num / 10
	res += num % 10
	num = num / 10
	res += num % 10
	num = num / 10
	res += num % 10
	num = num / 10
	fmt.Println(res)

	fmt.Println("Сумма цифр числа", num, "равна:", res)
}

func celsiusToFahrenheit() {
	var temp float64
	var unit string

	fmt.Print("Введите температуру и единицу измерения (C для Цельсия, F для Фаренгейта): ")
	fmt.Scan(&temp, &unit)

	switch unit {
	case "C", "c":
		// Цельсий в Фаренгейты
		fahrenheit := temp*9/5 + 32
		fmt.Printf("%.2f (Celsius) = %.2f (Fahrenheit)\n", temp, fahrenheit)
	case "F", "f":
		// Фаренгейты в Цельсий
		celsius := (temp - 32) * 5 / 9
		fmt.Printf("%.2f (Fahrenheit) = %.2f (Celsius)\n", temp, celsius)
	default:
		fmt.Println("Некорректная единица измерения. Используйте 'C' для Цельсия или 'F' для Фаренгейта.")
	}
}

func doubleArray() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	doubledArr := make([]int, len(arr))
	for i, v := range arr {
		doubledArr[i] = v * 2
	}

	fmt.Println("Удвоенный массив:", doubledArr)
}

func joinStrings() {
	var n int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&n)

	words := make([]string, n)
	fmt.Println("Введите строки:")
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}

	result := strings.Join(words, "")
	fmt.Println("Объединенная строка:", result)
}

func distance() {
	var x1, y1, x2, y2 float64
	fmt.Print("Введите координаты первой точки (x1, y1): ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Введите координаты второй точки (x2, y2): ")
	fmt.Scan(&x2, &y2)

	dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f) равно: %.2f\n", x1, y1, x2, y2, dist)
}

func isEven() {
	var number int
	fmt.Print("Введите число для проверки на четность: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Число", number, "четное")
	} else {
		fmt.Println("Число", number, "нечетное")
	}
}

func isLeapYear() {
	var year int
	fmt.Print("Введите год для проверки на високосный: ")
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Год", year, "високосный")
	} else {
		fmt.Println("Год", year, "невисокосный")
	}
}

func maxOfThree() {
	var a, b, c int
	fmt.Print("Введите три числа для нахождения наибольшего: ")
	fmt.Scan(&a, &b, &c)

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	fmt.Println("Наибольшее число из", a, b, c, "это", max)
}

func ageCategory() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	var category string
	if age < 13 {
		category = "Ребенок"
	} else if age < 18 {
		category = "Подросток"
	} else if age < 65 {
		category = "Взрослый"
	} else {
		category = "Пожилой"
	}

	fmt.Println("Возрастная категория:", category)
}

func isDivisibleBy3And5() {
	var number int
	fmt.Print("Введите число для проверки делимости на 3 и 5: ")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Число", number, "делится на 3 и 5")
	} else {
		fmt.Println("Число", number, "не делится на 3 и 5")
	}
}

func factorial() {
	var n int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&n)

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Факториал числа:", factorial)
}

func fibonacci() {
	var n int
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное значение, введите положительное число.")
		return
	}

	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
		for i := 2; i < n; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}

	fmt.Println("Первые", n, "чисел Фибоначчи:", fib)
}

func reverseArray() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	reversedArr := reverse(arr)
	fmt.Println("Перевернутый массив:", reversedArr)
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func primeNumbers() {
	var limit int
	fmt.Print("Введите предел для поиска простых чисел: ")
	fmt.Scan(&limit)

	primes := findPrimes(limit)
	fmt.Println("Простые числа до", limit, ":", primes)
}

func findPrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumArray() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sum := calculateSum(arr)
	fmt.Println("Сумма элементов массива:", sum)
}

func calculateSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
