package model

type Task struct{
	ID int
	Title string
	Description string
	Status string
}

type Tasks struct{
	tasks []Task
}