package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Sudoku struct {
	sudokuGrid      [9][9]int      //Array for Storing Values in cells
	replicatedGrid  [9][9]int      //Replica of Original Grid Before removing values
	answerGrid      [9][9]int      //Storing ANswers from Users
	difficultyLevel map[string]int //Map Storing K values based on level
	gridSize        int            //Size of Puzzle
	blockSize       int            //Size of a block
	gameLevel       string         //Game Level Chosen by Player
}

// Function for generating random value
func randomValueGenerator(upperLimit int) (randomInt int) {
	rand.Seed(time.Now().UnixNano())
	var rangeMin int = 1
	var rangeMax int = upperLimit + 1
	randomInt = rand.Intn(rangeMax-rangeMin) + rangeMin
	return
}

//Set Puzzle Settings from User
func (s *Sudoku) initializeGame(puzzleSize int, subBoxSize int, level string) {
	s.gridSize = puzzleSize
	s.blockSize = subBoxSize
	s.gameLevel = level
}

// Function to replicate a original grid
func replicateOriginalGrid(sudokugrid [9][9]int) [9][9]int {
	return sudokugrid
}

// Constant Values for Setting Levels
func (s *Sudoku) setKValue() {
	s.difficultyLevel = make(map[string]int)
	var key string
	for i := 0; i < 3; i++ {
		key = strconv.Itoa(i)
		s.difficultyLevel[key] = int(s.gridSize * 2 * (i + 1))
	}
}

// Functiom for Validating Uniqueness of Element in Row, Column and Box
func (s *Sudoku) uniqueValidation(rowNum int, colNum int, candidateValue int) bool {

	status := (s.uniqueRowValidation(rowNum, candidateValue) && s.uniqueColValidation(colNum, candidateValue) && s.uniqueBoxValidation(rowNum, colNum, candidateValue))
	return status
}

// Functiom for Validating Uniqueness of Element in a given Row
func (s *Sudoku) uniqueRowValidation(rowNum int, candidateValue int) bool {

	//check row wise
	for j := 0; j < s.gridSize; j++ {
		if s.sudokuGrid[rowNum][j] == candidateValue {
			return false
		}
	}
	return true
}

// Functiom for Validating Uniqueness of Element in a given Column
func (s *Sudoku) uniqueColValidation(colNum int, candidateValue int) bool {

	//check column wise
	for i := 0; i < s.gridSize; i++ {
		if s.sudokuGrid[i][colNum] == candidateValue {
			return false
		}
	}
	return true
}

func (s *Sudoku) uniqueBoxValidation(rowNum int, colNum int, candidateValue int) bool {

	rowStart := rowNum - (rowNum % s.blockSize)
	colStart := colNum - (colNum % s.blockSize)

	for i := 0; i < s.blockSize; i++ {
		for j := 0; j < s.blockSize; j++ {
			if s.sudokuGrid[rowStart+i][colStart+j] == candidateValue {
				return false
			}
		}
	}
	return true
}

// Function for Filling Cells in particular box in the Puzzle
func (s *Sudoku) fillIndividualBox(row, col int) {
	var num int
	for i := 0; i < s.blockSize; i++ {
		for j := 0; j < s.blockSize; j++ {
			for { // find unique number
				num = randomValueGenerator(s.gridSize)
				if s.uniqueBoxValidation(row, col, num) {
					break
				}
			}
			s.sudokuGrid[row+i][col+j] = num
		}
	}
}

// Function for Filling Cells the Diagonal boxes in the Puzzle
func (s *Sudoku) fillDiagonalBoxes() {
	for i := 0; i < s.gridSize; i += s.blockSize {
		s.fillIndividualBox(i, i)
	}
}

// Function for Filling Remaining Cells in the Puzzle i.e Other than Diagonal Boxes
func (s *Sudoku) fillRemainingCells(i int, j int) bool {

	if j >= s.gridSize && i < (s.gridSize-1) {
		i = i + 1
		j = 0
	}
	if i >= s.gridSize && j >= s.gridSize {
		return true
	}
	if i < s.blockSize {
		if j < s.blockSize {
			j = s.blockSize
		}
	} else if i < (s.gridSize - s.blockSize) {
		if j == int(i/s.blockSize)*s.blockSize {
			j = j + s.blockSize
		}
	} else {
		if j == (s.gridSize - s.blockSize) {
			i = i + 1
			j = 0
			if i >= s.gridSize {
				return true
			}
		}
	}

	for num := 1; num <= s.gridSize; num++ {
		if s.uniqueValidation(i, j, num) {
			s.sudokuGrid[i][j] = num
			if s.fillRemainingCells(i, j+1) {
				return true
			}
			s.sudokuGrid[i][j] = 0
		}
	}
	return false
}

//Function for removing K-values from grid based on difficulty Level
func (s *Sudoku) removeKCells(gameLevel string) {
	var count int = s.difficultyLevel[gameLevel]
	for {
		var cellID = randomValueGenerator(s.gridSize*s.gridSize - 1)

		i := (cellID / s.gridSize)
		j := (cellID % s.gridSize)

		if s.sudokuGrid[i][j] != 0 {
			count = count - 1
			s.sudokuGrid[i][j] = 0
		}
		if count == 0 {
			break
		}
	}
}

// Function for Generating Puzzle
func (s *Sudoku) createPuzzle(gameLevel string) {

	s.fillDiagonalBoxes()                                  //Fill the diagonal boxes
	s.fillRemainingCells(0, s.blockSize-1)                 //Fill remaining boxes except diagonal boxes
	s.replicatedGrid = replicateOriginalGrid(s.sudokuGrid) //Copy Unaltered Grid in New Grid
	s.setKValue()                                          //Set K values in map
	s.removeKCells(gameLevel)                              //Remove Cells form Grid Based on Difficulty of Game
}
