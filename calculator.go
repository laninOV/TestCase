package main

import "fmt"

func main() {
	fmt.Println("Какую операцию Вы бы хотели совершить?")
	fmt.Println("Выберите соответствующую цифру:")
	fmt.Println("1 - сложение")
	fmt.Println("2 - вычитание")
	fmt.Println("3 - умножение")
	fmt.Println("4 - деление")
	var choice, a, b int
	fmt.Scan(&choice)
	fmt.Println("С какими числами Вы бы хотели выполнить действие?")
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	fmt.Println("Полученный ответ:", calculator(a, b, choice))
}

func calculator(a, b, choice int) int {
	var result int

	switch choice {
	case 1:
		result = a + b
	case 2:
		result = a - b
	case 3:
		result = a * b
	case 4:
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Ошибка: деление на ноль")
			return 0
		}
	default:
		fmt.Println("Ошибка: неверный выбор операции")
		return 0
	}

	return result
}
