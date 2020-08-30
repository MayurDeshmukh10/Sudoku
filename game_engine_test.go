package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSetKValue(t *testing.T) {

	s := Sudoku{}
	s.gridSize = 9
	s.setKValue()
	answer := make(map[string]int)
	answer["0"] = 2 * 1 * 9
	answer["1"] = 2 * 2 * 9
	answer["2"] = 2 * 3 * 9

	if !reflect.DeepEqual(answer, s.difficultyLevel) {
		t.Error("Maps Didnot Match")
	}

}

func TestInitializeGame(t *testing.T) {

	s := Sudoku{}
	GridSize := 9
	BlockSize := 3
	Level := "0"
	s.initializeGame(GridSize, BlockSize, Level)

	if !assert.Equal(t, s.gridSize, GridSize) {
		t.Error("Grid Size didnot match")
	}

	if !assert.Equal(t, s.blockSize, BlockSize) {
		t.Error("Block Size didnot match")
	}

	if !assert.Equal(t, s.gameLevel, Level) {
		t.Error("Game Level didnot match")
	}
}

func TestRemoveKCells(t *testing.T) {
	s := Sudoku{}
	GridSize := 9
	BlockSize := 3
	Level := "0"
	s.initializeGame(GridSize, BlockSize, Level)
	s.createPuzzle(Level)

	count := 0

	for i := 0; i < s.gridSize; i++ {
		for j := 0; j < s.gridSize; j++ {
			if s.sudokuGrid[i][j] == 0 {
				count = count + 1
			}
		}
	}

	if !assert.Equal(t, count, s.difficultyLevel[s.gameLevel]) {
		t.Error("Number of Removed Cells not matched to the specification")
	}
}
