package cmd

import (
	"bufio"
	"fmt"
	"os"
	datastore "todo/pkg/datastore"
	model "todo/pkg/model"

	"github.com/spf13/cobra"
)

var status string

func init(){
	addCmd.Flags().StringVarP(&status,"status","s","Not Started",`Mark the status of a task
Status can be a String in List ("Not Started","In Progress","Completed") - Make sure to Add the Double Quotes ""
or you can use Short Hands for each possible status value (ns,ip,c)
By Default, It will be marked as Completed`)
	rootCmd.AddCommand(addCmd)
}

var addCmd=&cobra.Command{
	Use: "add",
	Short: "Add a new task",
	Long: `
Just Type the command "todo add" and it will prompt you for the required fields.
	`,
	Args: cobra.MaximumNArgs(0) ,
	Run: func(cmd *cobra.Command, args []string){
		task:=new(model.Task)
		for(task.Title==""){
		fmt.Println("Enter Title (Required) : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan(){
			task.Title=scanner.Text()
		}
		}
		fmt.Println("Enter Description ([Default] Empty) : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan(){
			task.Description=scanner.Text()
		}
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
		if err:=datastore.AddTask(*task);err!=nil{
			panic (err)
		}
	} ,
}