package main

import (
	"fmt"
	"os"
)

func check_win(type_player string) {
	var filled = 0
	for i := 0; i < 3; i++ {
		if boardArr[0][i] == "[X]" && boardArr[0][i] == boardArr[1][i] && boardArr[1][i] == boardArr[2][i] {
			fmt.Println("Игрок 1 победил")
			gameOver = true
			os.Exit(0)
		} else if boardArr[0][i] == "[0]" && boardArr[0][i] == boardArr[1][i] && boardArr[1][i] == boardArr[2][i] {
			if type_player == "bot" {
				fmt.Println("Бот победил")
			} else if type_player == "player2" {
				fmt.Println("Игрок 2 победил")
			}
			gameOver = true
			os.Exit(0)
		}

		if boardArr[i][0] == "[X]" && boardArr[i][0] == boardArr[i][1] && boardArr[i][1] == boardArr[i][2] {
			fmt.Println("Игрок 1 победил")
			gameOver = true
			os.Exit(0)
		} else if boardArr[i][0] == "[0]" && boardArr[i][0] == boardArr[i][1] && boardArr[i][1] == boardArr[i][2] {
			if type_player == "bot" {
				fmt.Println("Бот победил")
			} else if type_player == "player2" {
				fmt.Println("Игрок 2 победил")
			}
			gameOver = true
			os.Exit(0)
		}
	}

	if boardArr[0][0] == "[X]" && boardArr[0][0] == boardArr[1][1] && boardArr[1][1] == boardArr[2][2] {
		fmt.Println("Игрок 1 победил")
		gameOver = true
		os.Exit(0)
	} else if boardArr[0][0] == "[0]" && boardArr[0][0] == boardArr[1][1] && boardArr[1][1] == boardArr[2][2] {
		if type_player == "bot" {
			fmt.Println("Бот победил")
		} else if type_player == "player2" {
			fmt.Println("Игрок 2 победил")
		}
		gameOver = true
		os.Exit(0)
	}

	if boardArr[0][2] == "[X]" && boardArr[0][2] == boardArr[1][1] && boardArr[1][1] == boardArr[2][0] {
		fmt.Println("Игрок 1 победил")
		gameOver = true
		os.Exit(0)
	} else if boardArr[0][2] == "[0]" && boardArr[0][2] == boardArr[1][1] && boardArr[1][1] == boardArr[2][0] {
		if type_player == "bot" {
			fmt.Println("Бот победил")
		} else if type_player == "player2" {
			fmt.Println("Игрок 2 победил")
		}
		gameOver = true
		os.Exit(0)
	}

	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if boardArr[i][k] != "[ ]" {
				filled += 1
			}
		}
	}

	if filled == 9 {
		fmt.Println("Ничья")
		gameOver = true
		os.Exit(0)
	}
}
