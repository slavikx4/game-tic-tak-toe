package main

import (
	"errors"
	"fmt"
	"github.com/inancgumus/screen"
)

type Board struct {
	rows, cols int
	matrix     [][]string
}

func (board *Board) initialize() {
	for i := range board.matrix {
		board.matrix[i] = make([]string, board.cols)
		for j := range board.matrix[i] {
			board.matrix[i][j] = "."
		}
	}
}

func (board *Board) Print(res byte) {
	for _, el := range board.matrix {
		fmt.Println(el[0], " ", el[1], " ", el[2])
	}
	switch res {
	case 1:
		fmt.Print("ход игрока 1: ")
	case 2:
		fmt.Print("ход игрока 1: ")
	case 3:
		fmt.Print("ничья")
	case 4:
		fmt.Print("победил игрок 1")
	case 5:
		fmt.Print("победил игрок 2")
	}
}

func (board *Board) StepFirst() (bool, error) {
	var r, c byte
	_, err := fmt.Scan(&r, &c)
	if err != nil {
		panic(err)
	}
	if r > 3 || r < 1 || c > 3 || c < 1 {
		err := errors.New("невозможно так походить")
		return false, err
	}
	r-- //r - 1 : для пользовательского интерфейса
	c-- //c - 1 : для пользовательского интерфейса
	if board.matrix[r][c] == "." {
		board.matrix[r][c] = "X"
		if board.CheckWinFirst() {
			return true, nil
		}
	} else {
		err := errors.New("невозможно так походить")
		return false, err
	}
	return false, nil
}

func (board *Board) StepSecond() (bool, error) {
	var r, c byte
	_, err := fmt.Scan(&r, &c)
	if err != nil {
		panic(err)
	}
	if r > 3 || r < 1 || c > 3 || c < 1 {
		err := errors.New("невозможно так походить")
		return false, err
	}
	r-- //r - 1 : для пользовательского интерфейса
	c-- //c - 1 : для пользовательского интерфейса
	if board.matrix[r][c] == "." {
		board.matrix[r][c] = "O"
		if board.CheckWinSecond() {
			return true, nil
		}
	} else {
		err := errors.New("невозможно так походить")
		return false, err
	}
	return false, nil
}

func (board *Board) CheckWinFirst() bool {
	var win bool
	var cols0, cols1, cols2 string
	for _, el := range board.matrix {
		if el[0]+el[1]+el[2] == "XXX" {
			return true
		}
		cols0 += el[0]
		cols1 += el[1]
		cols2 += el[2]
	}
	switch {
	case string(cols0[0])+string(cols1[1])+string(cols2[2]) == "XXX":
		win = true
	case string(cols0[2])+string(cols1[1])+string(cols2[0]) == "XXX":
		win = true
	case string(cols0[0])+string(cols1[0])+string(cols2[0]) == "XXX":
		win = true
	case string(cols0[1])+string(cols1[1])+string(cols2[1]) == "XXX":
		win = true
	case string(cols0[2])+string(cols1[2])+string(cols2[2]) == "XXX":
		win = true
	}
	return win
}

func (board *Board) CheckWinSecond() bool {
	var win bool
	var cols0, cols1, cols2 string
	for _, el := range board.matrix {
		if el[0]+el[1]+el[2] == "OOO" {
			return true
		}
		cols0 += el[0]
		cols1 += el[1]
		cols2 += el[2]
	}
	switch {
	case string(cols0[0])+string(cols1[1])+string(cols2[2]) == "OOO":
		win = true
	case string(cols0[2])+string(cols1[1])+string(cols2[0]) == "OOO":
		win = true
	case string(cols0[0])+string(cols1[0])+string(cols2[0]) == "OOO":
		win = true
	case string(cols0[1])+string(cols1[1])+string(cols2[1]) == "OOO":
		win = true
	case string(cols0[2])+string(cols1[2])+string(cols2[2]) == "OOO":
		win = true
	}
	return win
}

func main() {
	// определяем размер доски
	rows := 3
	cols := 3
	//право на ход игрокам
	stepFirst, stepSecond := true, false
	// создаём доску
	board := Board{
		rows:   rows,
		cols:   cols,
		matrix: make([][]string, rows),
	}
	//инициализируем доску
	board.initialize()

	board.Print(1)

	//запуск игры
	for i := 0; i < rows*cols; {
		if stepFirst {
			win, err := board.StepFirst()
			if err != nil {
				fmt.Print(err.Error(), ": ")
			} else {
				if win {
					board.Print(4)
					break
				} else {
					i++
					if i == 9 {
						board.Print(3)
					} else {
						stepFirst, stepSecond = false, true
						screen.Clear()
						board.Print(2)
					}
				}
			}

		} else if stepSecond {
			win, err := board.StepSecond()
			if err != nil {
				fmt.Print(err.Error(), ": ")
			} else {
				if win {
					board.Print(5)
					break
				} else {
					i++
					if i == 9 {
						board.Print(3)
					} else {
						stepFirst, stepSecond = true, false
						screen.Clear()
						board.Print(1)
					}
				}
			}
		}
	}
}
