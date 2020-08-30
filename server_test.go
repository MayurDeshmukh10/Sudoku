package main

import(
	"fmt"
	"testing"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"time"
)

func assertJson(actual string, data []User, t *testing.T){
	jsonData, err:= json.Marshal(data)
	if err != nil{
		t.Fatalf("Error occured while json.Marshal: %s",err)
	}

	expected := string(jsonData)

	if actual != expected{
		t.Errorf("Expected result is different than actual")
	}
}

func TestShouldInsertRecord(t *testing.T){
	conn, mock, err := sqlmock.New()
	if err != nil{
		t.Fatalf("Connection Error: %s",err)
	}
	defer conn.Close()
	var duration time.Duration
	duration = 994324567892
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Leaderboard").WithArgs("Pikachu","0:16:34",0,9).WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	result := addPlayer(conn, "Pikachu", duration, 0, 9)
	
	if result != nil{
		t.Errorf("Error on Insert: %s", result)
  }
	
  expectations := mock.ExpectationsWereMet()
  if expectations != nil{
  	t.Errorf("Some expectations were not met: %s",expectations)
  }
}

func TestShouldNotInsertRecord(t *testing.T){
	conn, mock, err := sqlmock.New()
	if err != nil{
		t.Fatalf("Connection Error: %s",err)
	}
	defer conn.Close()

	var duration time.Duration
	duration = 994324567892
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Leaderboard").WithArgs("Pikachu","0:16:34",0,9).WillReturnError(fmt.Errorf("error on insertion"))
	mock.ExpectRollback()

	result := addPlayer(conn, "Pikachu", duration, 0, 9)
	
	if result == nil{
		t.Errorf("Expecting error on Insert: %s", result)
  }
	
  expectations := mock.ExpectationsWereMet()
  if expectations != nil{
  	t.Errorf("Some expectations were not met: %s",expectations)
  }	
}

func TestShouldFetchRecord(t *testing.T) {
	conn, mock, err := sqlmock.New()
	if err != nil{
		t.Fatalf("Connection Error: %s",err)
	}
	defer conn.Close()

	rows := sqlmock.NewRows([]string{"Name","Time"}).
		AddRow("Pikachu", "00:08:34").
		AddRow("Shinchan", "00:17:12")

	mock.ExpectQuery("SELECT (.+) FROM Leaderboard").WithArgs(0, 9).WillReturnRows(rows)

	result := getTopPlayers(conn, 0, 9)

	data := []User{	{Name:"Pikachu",Time:"00:08:34"},
										{Name:"Shinchan",Time:"00:17:12"},
								 }

	assertJson(result, data, t)

	err = mock.ExpectationsWereMet()
	if err != nil{
		t.Errorf("Expectations were not met %s",err)
	}
}
