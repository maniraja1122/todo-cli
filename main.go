package main

import (
	"todo/cmd"
	datastore "todo/pkg/datastore"
)

func main(){
	datastore.InitDB()
	cmd.Execute()
}