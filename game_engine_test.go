package main

import (
	"testing"

	"github.com/tj/assert"
)

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

// TestRandomValueGenerator function is used to check if randomValueGenerator is generating a number within the given range
func TestRandomValueGenerator(t *testing.T) {
	result := randomValueGenerator(9)
	assert.GreaterOrEqual(t, result, 1)
	assert.LessOrEqual(t, result, 9)
}

// TestReplicateOriginalGrid function is used to check is the replicateOriginalGrid is coping the grid properly
func TestReplicateOriginalGrid(t *testing.T) {
	a := [9][9]int{{1, 2, 3}, {4, 5, 6}}

	result := replicateOriginalGrid(a)
	assert.Equal(t, result, a)
}

// TestFillIndividualBox function is used to check if fillIndividualBox has filled all 1-9 numbers in the respective box with unique validation
func TestFillIndividualBox(t *testing.T) {
	// setup for test
	s := Sudoku{
		sudokuGrid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		gridSize:  9,
		blockSize: 3,
	}
	// row and col is the starting point of that respective individual box
	row := 3
	col := 3
	s.fillIndividualBox(row, col)
	check := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
		7: false,
		8: false,
		9: false,
	}
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			for num := 1; num <= 9; num++ {
				// check if all 1-9 numbers are filled in the respective diagonal box
				if num == s.sudokuGrid[i][j] {
					check[num] = true
					break
				}
			}
		}
	}
	for _, v := range check {
		// check if any number between 1-9 is not filled in the respective diagonal box
		if !v {
			t.Fatalf("FillIndividualBox function has failed to fill all 1-9 number in box")
			break
		}
	}
}

// TestFillDiagonalBox function is used for testing whether fillDiagonalBox function has filled all the cells in diagonal box with unique validation
func TestFillDiagonalBox(t *testing.T) {
	// setup for test
	s := Sudoku{
		sudokuGrid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		gridSize:  9,
		blockSize: 3,
	}
	s.fillDiagonalBoxes()
	// to check each diagonal box in grid
	for i := 0; i < s.blockSize; i++ {
		flag := false
		row := 3 * i
		col := 3 * i
		check := map[int]bool{
			1: false,
			2: false,
			3: false,
			4: false,
			5: false,
			6: false,
			7: false,
			8: false,
			9: false,
		}
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				for num := 1; num <= 9; num++ {
					// check if all 1-9 numbers are filled in the respective diagonal box
					if num == s.sudokuGrid[i][j] {
						check[num] = true
						break
					}
				}
			}
		}
		for _, v := range check {
			// check if any number between 1-9 is not filled in the respective diagonal box
			if !v {
				flag = true
				break
			}
		}
		if flag {
			t.Fatalf("FillIndividualBox function has failed to fill all 1-9 number in box")
			break
		}
	}
}

// checkdiagonalboxes function is used to check if the cell is in diagonal box
func checkdiagonalboxes(i, j, blockSize int) bool {
	if (i >= (blockSize*(j/blockSize)) && j >= (blockSize*(i/blockSize))) && (i < (blockSize*(j/blockSize+1)) && j < (blockSize*(i/blockSize+1))) {
		return true
	}
	return false
}

// TestFillRemainingCells function is used for testing whether fillRemainingCell function has filled all the remaining cells( i.e. cells except diagonal box cells) in the grid with unique validation
func TestFillRemainingCells(t *testing.T) {
	// setup for test
	s := Sudoku{
		sudokuGrid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		gridSize:  9,
		blockSize: 3,
	}
	// s.fillDiagonalBoxes()
	s.fillRemainingCells(0, s.blockSize)
	flag := false
	for i := 0; i < s.gridSize; i++ {
		for j := 0; j < s.gridSize; j++ {
			// skipping diagonal box cells
			if checkdiagonalboxes(i, j, s.blockSize) {
				continue
			}
			// check is fillRemainingCells function has filled all remaining cells( i.e. cells except diagonal box cells) and unique validation
			if 0 == s.sudokuGrid[i][j] || s.uniqueValidation(i, j, s.sudokuGrid[i][j]) {
				flag = true
				break
			}
		}
	}
	if flag {
		t.Fatalf("FillIndividualBox function has failed to fill all 1-9 number in box")
	}
}
