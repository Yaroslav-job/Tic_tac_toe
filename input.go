package main

import (
	"fmt"
)

func input() (int, int) {
	var column, line int
	fmt.Println("Введите номер колонки и номер строки на пересечении которых находится")
	for {
		fmt.Print("Номер колонки:")
		fmt.Scan(&column)
		fmt.Print("Номер строки:")
		fmt.Scan(&line)
		if column >= 1 && column <= 3 && line >= 1 && line <= 3 {
			break
		} else if column < 1 || column > 3 {
			fmt.Println("Вы ввели несуществующий номер колонки")
		} else if line < 1 || line > 3 {
			fmt.Println("Вы ввели несуществующий номер строки")
		}
	}
	return column, line
}

func bot() (int, int) {
	for {
		column_bot := random()
		line_bot := random()
		if boardArr[line_bot-1][column_bot-1] == "[X]" || boardArr[line_bot-1][column_bot-1] == "[0]" {
			fmt.Println("Этот ход не возможен. Ячейка уже занята")
		} else {
			return column_bot, line_bot
		}
	}
}
