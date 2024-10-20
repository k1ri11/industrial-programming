package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
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

// 1. Проверка на простоту
func task1_1() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if isPrime(num) {
		fmt.Println(num, "простое число.")
	} else {
		fmt.Println(num, "не является простым числом.")
	}
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

// 2. Нахождение НОД (алгоритм Евклида)
func task1_2() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	fmt.Println("НОД:", gcd(a, b))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 3. Сортировка пузырьком
func task1_3() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 4. Таблица умножения
func task1_4() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// 5. Фибоначчи с мемоизацией
var fibCache = map[int]int{}

func task1_5() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Println("Число Фибоначчи:", fibonacciMemo(n))
}

func fibonacciMemo(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := fibCache[n]; ok {
		return val
	}
	fibCache[n] = fibonacciMemo(n-1) + fibonacciMemo(n-2)
	return fibCache[n]
}

// 6. Обратные числа
func task2_1() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Println("Число в обратном порядке:", reverseNumber(num))
}

func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

// 7. Треугольник Паскаля
func task2_2() {
	var rows int
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)
	printPascalsTriangle(rows)
}

func printPascalsTriangle(n int) {
	triangle := make([][]int, n)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		fmt.Println(triangle[i])
	}
}

// 8. Число палиндром (без использования строк)
func task2_3() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if isPalindromeNumber(num) {
		fmt.Println("Число палиндром.")
	} else {
		fmt.Println("Число не является палиндромом.")
	}
}

func isPalindromeNumber(num int) bool {
	original := num
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return original == reversed
}

// 9. Максимум и минимум в массиве
func task2_4() {
	arr := []int{12, 3, 5, 7, 19, 0, -3}
	min, max := findMinMax(arr)
	fmt.Println("Минимум:", min, "Максимум:", max)
}

func findMinMax(arr []int) (min, max int) {
	min, max = arr[0], arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return
}

// 10. Игра "Угадай число"
func task2_5() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess, attempts int
	for attempts < 10 {
		fmt.Print("Введите ваше предположение: ")
		fmt.Scan(&guess)
		attempts++
		if guess < target {
			fmt.Println("Больше")
		} else if guess > target {
			fmt.Println("Меньше")
		} else {
			fmt.Println("Поздравляем! Вы угадали.")
			return
		}
	}
	fmt.Println("Извините, вы не угадали. Число было:", target)
}

// 11. Числа Армстронга
func task3_1() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if isArmstrong(num) {
		fmt.Println(num, "число Армстронга.")
	} else {
		fmt.Println(num, "не является числом Армстронга.")
	}
}

func isArmstrong(num int) bool {
	sum, temp, digits := 0, num, int(math.Log10(float64(num)))+1
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == num
}

// 12. Подсчет уникальных слов в строке
func task3_2() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите предложение:")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)
	uniqueWords := countUniqueWords(sentence)
	fmt.Println("Количество уникальных слов:", uniqueWords)
}

func countUniqueWords(s string) int {
	words := strings.Fields(s)
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}
	return len(wordMap)
}

// 13. Игра "Жизнь"
func task3_3() {
	fmt.Println("Игра Жизнь:")
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for i := 0; i < 50; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a // Swap universes
	}

}

const (
	width  = 10
	height = 10
)

// Universe является двухмерным полем клеток.
type Universe [][]bool

// NewUniverse возвращает пустую вселенную.
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Seed заполняет вселенную случайными живыми клетками.
func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

// Set устанавливает состояние конкретной клетки.
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// Alive сообщает, является ли клетка живой.
// Если координаты за пределами вселенной, возвращаемся к началу.
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// Neighbors подсчитывает прилегающие живые клетки.
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

// Next возвращает состояние определенной клетки на следующем шаге.
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// String возвращает вселенную как строку
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}

	return string(buf)
}

// Show очищает экран и возвращает вселенную.
func (u Universe) Show() {
	fmt.Print(u.String())
}

// Step обновляет состояние следующей вселенной (b) из
// текущей вселенной (a).
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

// 14. Цифровой корень числа
func task3_4() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Println("Цифровой корень числа:", digitalRoot(num))
}

func digitalRoot(num int) int {
	for num > 9 {
		num = sumDigits(num)
	}
	return num
}

func sumDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// 15. Преобразование арабского числа в римское
func task3_5() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Println("Римское число:", intToRoman(num))
}

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			result += symbols[i]
			num -= vals[i]
		}
	}
	return result
}
