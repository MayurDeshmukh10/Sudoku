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
    	return true
}
  
func (s *Sudoku) uniqueBoxValidation(rowNum int, colNum int, candidateValue int) bool {

	// Getting index of first element of that block
	rowStart := (rowNum / s.blockSize) * s.blockSize
	colStart := (colNum / s.blockSize) * s.blockSize

	for i := 0; i < s.blockSize; i++ {
		for j := 0; j < s.blockSize; j++ {
			if s.sudokuGrid[rowStart+i][colStart+j] == candidateValue {
				return false
			}
		}
	}
  return true
}


func main() {

}
