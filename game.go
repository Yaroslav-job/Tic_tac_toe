package main

import (
	"fmt"
)

var gameOver bool

func play_game(type_player string) {
	for !gameOver {
		column, line := input()
		if boardArr[line-1][column-1] == "[X]" || boardArr[line-1][column-1] == "[0]" {
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
			if boardArr[line-1][column-1] == "[X]" || boardArr[line-1][column-1] == "[0]" {
				fmt.Println("Этот ход не возможен. Ячейка уже занята")
				continue
			} else {
				clear()
				board(column, line, "player2")
			}
		}

	}
}
