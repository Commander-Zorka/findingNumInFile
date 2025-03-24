package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()

	// Увеличиваем буфер сканера для длинных строк
	const maxCapacity = 10 * 1024 * 1024 // 10 МБ
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	data := scanner.Text()

	// Разбиваем строку с учетом возможных пробелов
	numbers := strings.Split(strings.TrimSpace(data), ";")

	for pos, num := range numbers {
		num = strings.TrimSpace(num) // Удаляем возможные пробелы вокруг числа
		if num == "0" {
			fmt.Printf("Число 0 найдено на позиции: %d\n", pos+1)
			return
		}
	}

	fmt.Println("Число 0 не найдено в файле")

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения: %v\n", err)
	}
}
