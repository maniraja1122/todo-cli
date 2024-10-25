package datastore

import (
	"fmt"
	model "todo/pkg/model"
)

func AddTask(task model.Task) error{
	_,err:=DB.Exec("INSERT INTO tasks(title,description,status) VALUES(?,?,?)",task.Title,task.Description,task.Status)
	if err!=nil{
		return err
	}
	fmt.Println("New Task Added")
	return nil
}

func ListAllTasks() error{
	println("ID     Task")
	rows,err:=DB.Query("SELECT * FROM tasks")
	if err!=nil{
		return err
	}
	for rows.Next(){
		task:=model.Task{}
		if err:=rows.Scan(&task.ID,&task.Title,&task.Description,&task.Status); err!=nil{
			return err
		}
		fmt.Println(fmt.Sprintf("%d     %s",task.ID,task.Title))
	}
	return nil
}

func ListTaskbyStatus(status string) error{
	println("ID     Task")
	rows,err:=DB.Query("SELECT * FROM tasks WHERE status=?",status)
	if err!=nil{
		return err
	}
	for rows.Next(){
		task:=model.Task{}
		if err:=rows.Scan(&task.ID,&task.Title,&task.Description,&task.Status); err!=nil{
			return err
		}
		fmt.Println(fmt.Sprintf("%d     %s",task.ID,task.Title))
	}
	return nil
}

func DeleteTasks(id int) error{
	_,err:=DB.Exec("DELETE FROM tasks WHERE id=?",id)
	if err!=nil{
		return err
	}
	println("Task Deleted Successfully")
	return nil
}

func EditTasks(id int,task model.Task) error{
	_,err:=DB.Exec("UPDATE tasks SET title=?,description=?,status=? WHERE id=?",task.Title,task.Description,task.Status,id)
	if err!=nil{
		return err
	}
	return nil
}

func MarkTask(id int, status string) error{
	_,err:=DB.Exec("UPDATE tasks SET status=? WHERE id=?",status,id)
	if err!=nil{
		return err
	}
	return nil
}

func ShowTask(id int) error{
	rows,err:=DB.Query("Select * FROM tasks WHERE id=?",id)
	if err!=nil{
		return err
	}
	for rows.Next(){
		task:=model.Task{}
		if err:=rows.Scan(&task.ID,&task.Title,&task.Description,&task.Status);err!=nil{
			return err
		}
		println(fmt.Sprintf("Title : %s",task.Title))
		println(fmt.Sprintf("Description : %s",task.Description))
		println(fmt.Sprintf("Status : %s",task.Status))
	}
	return nil
}

func GetTask(id int) (task *model.Task,err error){
	rows,err:=DB.Query("Select * FROM tasks WHERE id=?",id)
	if err!=nil{
		return nil,err
	}
	task=new(model.Task)
	if rows.Next(){
		if err:=rows.Scan(&task.ID,&task.Title,&task.Description,&task.Status);err!=nil{
			return nil,err
		}
	} else {
		println("No Task Exists with that ID")
		return nil,nil
	}
	return task,nil
}