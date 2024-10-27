package datastore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB


func InitDB() error{
	const create string = `
  CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY,
  title TEXT(255) NOT NULL,
  description TEXT(10000),
  status TEXT(255) NOT NULL
  )`
  	var err error
	DB,err=sql.Open("sqlite3","./tasks.db")
	if err!=nil{
		return err
	}
	_,err=DB.Exec(create)
	if err!=nil{
		return err
	}
	return nil
}