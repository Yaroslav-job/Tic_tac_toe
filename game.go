package main

import (
	"fmt"
)

var gameOver bool

func play_game() {
	for !gameOver {
		column, line := input()
		if boardArr[line-1][column-1] == "[X]" || boardArr[line-1][column-1] == "[0]" {
			fmt.Println("Этот ход не возможен. Ячейка уже занята")
		} else {
			clear()
			board(column, line, "player")
		}
		column_bot, line_bot := bot()
		clear()
		board(column_bot, line_bot, "bot")
	}
}

func check_win() {
	for i := 0; i < 3; i++ {
		if boardArr[0][i] == "[X]" && boardArr[0][i] == boardArr[1][i] && boardArr[1][i] == boardArr[2][i] {
			fmt.Println("Игрок победил")
			gameOver = true
		} else if boardArr[0][i] == "[0]" && boardArr[0][i] == boardArr[1][i] && boardArr[1][i] == boardArr[2][i] {
			fmt.Println("Бот победил")
			gameOver = true
		}

		if boardArr[i][0] == "[X]" && boardArr[i][0] == boardArr[i][1] && boardArr[i][1] == boardArr[i][2] {
			fmt.Println("Игрок победил")
			gameOver = true
		} else if boardArr[i][0] == "[0]" && boardArr[i][0] == boardArr[i][1] && boardArr[i][1] == boardArr[i][2] {
			fmt.Println("Бот победил")
			gameOver = true
		}
	}

	if boardArr[0][0] == "[X]" && boardArr[0][0] == boardArr[1][1] && boardArr[1][1] == boardArr[2][2] {
		fmt.Println("Игрок победил")
		gameOver = true
	} else if boardArr[0][0] == "[0]" && boardArr[0][0] == boardArr[1][1] && boardArr[1][1] == boardArr[2][2] {
		fmt.Println("Бот победил")
		gameOver = true
	}

	if boardArr[0][2] == "[X]" && boardArr[0][2] == boardArr[1][1] && boardArr[1][1] == boardArr[2][0] {
		fmt.Println("Игрок победил")
		gameOver = true
	} else if boardArr[0][2] == "[0]" && boardArr[0][2] == boardArr[1][1] && boardArr[1][1] == boardArr[2][0] {
		fmt.Println("Бот победил")
		gameOver = true
	}
}
