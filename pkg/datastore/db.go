package datastore

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB


func InitDB() error{
	const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  title VARCHAR(255) UNIQUE NOT NULL PRIMARY KEY,
  description TEXT(10000),
  status ENUM("Not Started","In Progress","Completed") NOT NULL
  );`
	file, err := os.Create("todo-database.db")
	file.Close()
	if err!=nil{
		return err
	}
	DB,err=sql.Open("sqlite3","./todo-database.db")
	if err!=nil{
		return err
	}
	_,err=DB.Query(create)
	if err!=nil{
		return err
	}
	return nil
}