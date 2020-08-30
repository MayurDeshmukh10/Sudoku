package main

import "testing"

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

	// generating start indexof block
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
	value := s.uniqueValidation(rowStart, colStart, candidateValue)
	if value == false {
		//fmt.Println(s.sudokuGrid)
		t.Errorf("Duplicate didn't exist, should return true")
	}

	// for explantion of this see TestUniqueColValidationValid, TestUniqueRowValidationValid and TestUniqueBoxValidationValid functions

}
