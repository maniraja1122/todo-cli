package main

import (
	"todo/cmd"
	datastore "todo/pkg/datastore"
)
func init(){
	if err:=datastore.InitDB();err!=nil{
		panic(err)
	}
}

func main(){
	cmd.Execute()
}