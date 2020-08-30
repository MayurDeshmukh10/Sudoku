package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func createSudoku() *Sudoku {
	s := Sudoku{}
	// gameLevel is zero for now
	s.initializeGame(9, 3, "0")
	s.createPuzzle("0")

	return &s
}

//Test For  setKValue Function
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

//Test For initalizeGame Function
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

//Test For removeKCells-Success Function
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

//Test For removeKCells-Failure Function
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

func TestUniqueRowValidationValid(t *testing.T) {

	/*
		Logic For Valid Unique Row Testing
		We will find any entry which is 0
		Then we will fetch its index and store it in rowNum,colNum
		Then we will find any element which is absent in rowNum
		We will then validate this row with that element
	*/

	s := createSudoku()
	//comman function upload by santosh remove these comment at the end

	// Find first 0 entry's index
	rowNum, colNum := 0, 0

Loop:
	for rowNum = 0; rowNum < 9; rowNum++ {
		for colNum = 0; colNum < 9; colNum++ {
			if s.sudokuGrid[rowNum][colNum] == 0 {
				break Loop
			}
		}
	}

	checkNumPresent := [10]bool{}
	candidateValue := 0

	// Used to mark numbers present in the row
	for i := 0; i < 9; i++ {
		checkNumPresent[s.sudokuGrid[rowNum][i]] = true
	}

	// check which value can be a Candidate in inserting into valid row
	for i := 1; i < 10; i++ {
		if checkNumPresent[i] == false {
			candidateValue = i
			break
		}
	}

	// Validate
	value := s.uniqueRowValidation(rowNum, candidateValue)
	if value == false {
		t.Errorf("Duplicate didn't exist, should return true")
	}
	/*
		Example to trace above Test Case:

		Suppose Row Number 3 is considered as below:
		4  0  8  0  2  5  9  3  7

		Then RowNum= 3
		Then we know that at index(3,1) element present is 0
		We then check which value can be present here
		We get 1 as candidateValue
		Then we validate with s.uniqueRowValidation(rowNum, candidateValue)
		And as it is unique, value will become true
	*/

}

func TestUniqueColValidationValid(t *testing.T) {
	/*
		Logic For Valid Unique Column Testing
		We will find any entry which is 0
		Then we will fetch its index and store it in rowNum, colNum
		Then we will find any element which is absent in colNum
		we will then validate this column with that element
	*/

	s := createSudoku()

	// Find first 0 entry's index
	rowNum, colNum := 0, 0
Loop:
	for rowNum = 0; rowNum < 9; rowNum++ {
		for colNum = 0; colNum < 9; colNum++ {
			if s.sudokuGrid[rowNum][colNum] == 0 {
				break Loop
			}
		}
	}

	checkNumPresent := [10]bool{}
	candidateValue := 0

	// Used to mark numbers present in the col
	for i := 0; i < 9; i++ {
		checkNumPresent[s.sudokuGrid[i][colNum]] = true
	}

	// check which value can be a Candidate in inserting into valid col
	for i := 1; i < 10; i++ {
		if checkNumPresent[i] == false {
			candidateValue = i
			break
		}
	}

	// Validate
	value := s.uniqueColValidation(colNum, candidateValue)
	if value == false {
		t.Errorf("Duplicate didn't existed, should return true")
	}
	/*
		Example to trace above Test Case:

		Suppose Col Number 3 is considered as below:
		4
		0
		8
		0
		2
		5
		9
		3
		7

		Then ColNum= 3
		Then we know that at index(1,3) element present is 0
		We then check which value can be present here
		We get 1 as candidateValue
		Then we validate with s.uniqueRowValidation(rowNum, candidateValue)
		And as it is unique, value will become true
	*/
}

func TestUniqueBlockValidationValid(t *testing.T) {
	/*
		Logic For Valid Unique Block Testing
		We will find any entry which is 0
		Then we will fetch its index and store it in rowNum,colNum
		Then we will find any element which is absent in the block
		We will then validate this block with that element
	*/

	s := createSudoku()

	// Find first 0 entry's index
	rowNum, colNum := 0, 0
Loop:
	for rowNum = 0; rowNum < 9; rowNum++ {
		for colNum = 0; colNum < 9; colNum++ {
			if s.sudokuGrid[rowNum][colNum] == 0 {
				break Loop
			}
		}
	}

	checkNumPresent := [10]bool{}
	candidateValue := 0

	// generating start indexof block
	rowStart := (rowNum / 3) * 3
	colStart := (colNum / 3) * 3

	// Used to mark numbers present in the block
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			checkNumPresent[s.sudokuGrid[i][j]] = true
		}
	}

	// check which value can be a Candidate in inserting into valid block
	for i := 1; i < 10; i++ {
		if checkNumPresent[i] == false {
			candidateValue = i
			break
		}
	}

	// Validate
	value := s.uniqueBoxValidation(rowStart, colStart, candidateValue)
	if value == false {
		t.Errorf("Duplicate didn't exist, should return true")
	}
	/*
		Example to trace above Test Case:

		Suppose Row Number 0 , colNum 1 is considered as below:
		4  0  8
		0  2  5
		9  3  7

		Then we know that at index(0,1) element present is 0
		So, rowNum = 0 and  colNum = 1
		We then check which value can be present here
		We get 1 as candidateValue
		Then we validate with s.uniqueBoxValidation(rowStart, colStart, candidateValue)
		And as it is unique, value will become true
	*/
}

func TestUniqueValidationValid(t *testing.T) {
	/*
		Logic For Valid Unique Validation Testing
		We will find any entry which is 0
		Then we will fetch its index and store it in rowNum,colNum
		Then we will find any element which is absent in rowNum, colNum, as well as its block
		We will then validate this row with that element
	*/

	s := createSudoku()

	// Find first 0 entry's index
	rowNum, colNum := 0, 0
Loop:
	for rowNum = 0; rowNum < 9; rowNum++ {
		for colNum = 0; colNum < 9; colNum++ {
			if s.sudokuGrid[rowNum][colNum] == 0 {
				break Loop
			}
		}
	}

	checkNumPresent := [10]bool{}
	candidateValue := 0

	// generating start index of block
	rowStart := (rowNum / 3) * 3
	colStart := (colNum / 3) * 3

	// Used to mark numbers present in the row
	for i := 0; i < 9; i++ {
		checkNumPresent[s.sudokuGrid[rowNum][i]] = true
	}

	// Used to mark numbers present in the col
	for i := 0; i < 9; i++ {
		checkNumPresent[s.sudokuGrid[i][colNum]] = true
	}

	// Used to mark numbers present in the block
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			checkNumPresent[s.sudokuGrid[i][j]] = true
		}
	}

	// check which value can be a Candidate in inserting into valid block, row, col combination
	for i := 1; i < 10; i++ {
		if checkNumPresent[i] == false {
			candidateValue = i
			break
		}
	}

	// Validate
	value := s.uniqueValidation(rowNum, colNum, candidateValue)
	if value == false {
		t.Errorf("Duplicate didn't exist, should return true")
	}
}
