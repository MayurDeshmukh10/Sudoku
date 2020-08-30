# Sudoku-Game

### Dependencies

```
go get -u github.com/go-sql-driver/mysql
go get -u github.com/gorilla/mux
go get github.com/gorilla/websocket
go get github.com/urfave/negroni
go get github.com/stretchr/testify/assert

```

### Import MySQL Database

Set up: Create the database in your local database.

```
CREATE DATABASE sudoku;
mysql -u [username] -p sudoku < sudoku.sql
```

### Setting Enviroment Variables

```
Add following in bashrc file
export DATABASE_USERNAME="Your database username"
export DATABASE_PASSWORD="Your database password"
export DATABASE_NAME="Your database name"
```

### To run

```
go build
./Sudoku

Go to - localhost:3000
```

### How to run testcases

To run testcases:

```
go test
```

To run testcases and generate a html coverage report

```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```
