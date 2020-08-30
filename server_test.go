package main

import(
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"time"
)

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
		t.Fatalf("Error on Insert: %s", result)
  }
	
  expectations := mock.ExpectationsWereMet()
  if expectations != nil{
  	t.Fatalf("Some expectations were not met: %s",expectations)
  }
}