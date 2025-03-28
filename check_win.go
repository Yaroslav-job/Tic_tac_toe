package main

import (
	"fmt"
)

func check_win(type_player string) {
	var filled = 0
	for i := 0; i < 3; i++ {
		if (boardArr[0][i] == "[❌]" || boardArr[0][i] == "[⭕]") && boardArr[0][i] == boardArr[1][i] && boardArr[1][i] == boardArr[2][i] {
			fmt.Printf("Победил %s", type_player)
			game_over()
		}

		if (boardArr[i][0] == "[❌]" || boardArr[i][0] == "[⭕]") && boardArr[i][0] == boardArr[i][1] && boardArr[i][1] == boardArr[i][2] {
			fmt.Printf("Победил %s", type_player)
			game_over()
		}
	}

	if (boardArr[0][0] == "[❌]" || boardArr[0][0] == "[⭕]") && boardArr[0][0] == boardArr[1][1] && boardArr[1][1] == boardArr[2][2] {
		fmt.Printf("Победил %s", type_player)
		game_over()
	}

	if (boardArr[0][2] == "[❌]" || boardArr[0][2] == "[⭕]") && boardArr[0][2] == boardArr[1][1] && boardArr[1][1] == boardArr[2][0] {
		fmt.Printf("Победил %s", type_player)
		game_over()
	}

	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if boardArr[i][k] != "[  ]" {
				filled += 1
			}
		}
	}

	if filled == 9 {
		fmt.Println("Ничья")
		game_over()
	}
}
