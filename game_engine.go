package main

import (
	"math/rand"
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

//Set Puzzle Settings from User
func (s *Sudoku) initializeGame(puzzleSize int, subBoxSize int, level string) {
	s.gridSize = puzzleSize
	s.blockSize = subBoxSize
	s.gameLevel = level
}

func main() {

}
