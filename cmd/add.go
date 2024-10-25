package cmd

import (
	"fmt"
	model "todo/pkg/model"

	"github.com/spf13/cobra"
)

var status string

func init(){
	addCmd.Flags().StringVarP(&status,"status","s","Not Started","Set the status of the task")
	rootCmd.AddCommand(addCmd)
}

var addCmd=&cobra.Command{
	Use: "add",
	Short: "Add a new task",
	Args: cobra.MaximumNArgs(0) ,
	Run: func(cmd *cobra.Command, args []string){
		task:=new(model.Task)
		fmt.Println("Enter Title (Required) : ")
		fmt.Scanln(&task.Title)
		fmt.Println("Enter Description ([Default] Empty) : ")
		fmt.Scanln(&task.Description)
		fmt.Println("Status should a [String] in List (Not Started,In Progress,Completed) or Short Hands (ns,ip,c) : ")
		fmt.Scanln(&task.Description)
		switch status{
		case "Not Started","ns":
			task.Status="Not Started"
		case "In Progress","ip":
			task.Status="In Progress"
		case "Completed","c":
			task.Status="Completed"
		default:
			task.Status="Not Started"
		}
	} ,
}