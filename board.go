package main

import (
	"fmt"
)

var boardArr [3][3]string

func board_start() {
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			boardArr[i][k] = "[  ]"
		}
	}
	print_board()
}

func print_board() {
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			fmt.Printf("%s", boardArr[i][k])
		}
		fmt.Print("\n")
	}
}

func board(column int, line int, type_player string) {
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if i == line-1 && k == column-1 && type_player == "player" {
				boardArr[i][k] = "[❌]"
			} else if i == line-1 && k == column-1 && type_player == "bot" || i == line-1 && k == column-1 && type_player == "player2" {
				boardArr[i][k] = "[⭕]"
			}
		}
	}
	print_board()
	check_win(type_player)
}
