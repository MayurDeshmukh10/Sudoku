package main

import "testing"

func TestRemoveKCellsValid(t *testing.T) {
	s := Sudoku{}
	gameLevel := "0"
	s.initializeGame(9, 3, gameLevel)

	s.createPuzzle(gameLevel)

	var count int = s.difficultyLevel[gameLevel]

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.sudokuGrid[i][j] == 0 {
				count--
			}
		}
	}

	if count != 0 {
		t.Errorf("K cells weren't removed!")
	}
}

func TestRemoveKCellsInvalid(t *testing.T) {
	s := Sudoku{}
	// using different gameLevel, thus different count of blanks
	gameLevel := "1"
	s.initializeGame(9, 3, gameLevel)

	s.createPuzzle(gameLevel)

	var count int = s.difficultyLevel[gameLevel]

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.sudokuGrid[i][j] == 0 {
				count--
			}
		}
	}
	count++

	if count == 0 {
		t.Errorf("K cells were removed! That is wrong!")
	}
}
