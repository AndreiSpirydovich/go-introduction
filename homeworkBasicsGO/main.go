package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Делить на ноль нельзя")
	}
	return a / b, nil
}

func main() {
	var a, b int
	var ans string

	for {
		fmt.Println("Введите делимое")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Ошибка ввода, введите ЦЕЛОЕ число")
			return
		}

		fmt.Println("Введите делитель")
		_, err = fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Ошибка ввода, введите ЦЕЛОЕ число")
			return
		}

		result, err := divide(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println("Частное:", result)
		switch {
		case result > 10:
			fmt.Println("Результат большой")
		case result >= 1 && result <= 10:
			fmt.Println("Результат средний")
		case result <= 0:
			fmt.Println("Результат маленький или ноль")
		}

		fmt.Println("Хотите повторить?")
		fmt.Println("y - да, n - нет")
		fmt.Scanln(&ans)
		if ans == "n" {
			fmt.Println("Завершение")
			break
		} else {
			continue
		}
	}
}
