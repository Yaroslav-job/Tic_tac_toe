package main

import (
	"fmt"
	"os"
)

func play_game(type_player string) {
	for {
		column, line := input()
		if boardArr[line-1][column-1] == "[❌]" || boardArr[line-1][column-1] == "[⭕]" {
			fmt.Println("Этот ход не возможен. Ячейка уже занята")
			continue
		} else {
			clear()
			board(column, line, "player")
		}

		if type_player == "bot" {
			column_bot, line_bot := bot()
			clear()
			board(column_bot, line_bot, "bot")
		} else if type_player == "player" {
			column, line := input()
			if boardArr[line-1][column-1] == "[❌]" || boardArr[line-1][column-1] == "[⭕]" {
				fmt.Println("Этот ход не возможен. Ячейка уже занята")
				continue
			} else {
				clear()
				board(column, line, "player2")
			}
		}

	}
}

func game_over() {
	var expectation int
	fmt.Println("\n\nВведите: \n1 - чтобы начать игру заново \n2 - чтобы закончить игру")
	for {
		fmt.Scan(&expectation)
		if expectation == 1 {
			main()
		} else if expectation == 2 {
			os.Exit(0)
		} else {
			fmt.Printf("Вы ввели недопустимое значение. Попробуйте ещё раз:")
		}
	}
}
