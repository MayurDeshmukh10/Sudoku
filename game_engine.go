package main

type Sudoku struct {
	sudokuGrid      [9][9]int      //Array for Storing Values in cells
	replicatedGrid  [9][9]int      //Replica of Original Grid Before removing values
	answerGrid      [9][9]int      //Storing ANswers from Users
	difficultyLevel map[string]int //Map Storing K values based on level
	gridSize        int            //Size of Puzzle
	blockSize       int            //Size of a block
	gameLevel       string         //Game Level Chosen by Player
}

func main() {

}
