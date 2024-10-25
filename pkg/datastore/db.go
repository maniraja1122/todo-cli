package datastore

import (
	"database/sql"
	"fmt"
	// "os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)
const (
	host     = "fiberdb-demo.cfguaag2yn2c.us-west-2.rds.amazonaws.com"
	port     = 3306 // Default port
	user     = "admin"
	password = "password"
	dbname   = "fiberdb"
)
var DB *sql.DB


func InitDB() error{
	const create string = `
  CREATE TABLE IF NOT EXISTS tasks (
  id INT NOT NULL PRIMARY KEY,
  title TEXT(255) NOT NULL,
  description TEXT(10000),
  status TEXT(255) NOT NULL
  )`
	// file, err := os.Create("todo-database.db")
	// file.Close()
	// if err!=nil{
	// 	return err
	// }
	var err error=nil
	DB,err=sql.Open("mysql",fmt.Sprintf("%s:%s@(%s:%d)/%s",user,password,host,port,dbname))
	if err!=nil{
		return err
	}
	_,err=DB.Query(create)
	if err!=nil{
		return err
	}
	return nil
}