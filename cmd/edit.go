package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"fmt"
	"bufio"
	"os"
	datastore "todo/pkg/datastore"
)

func init(){
	rootCmd.AddCommand(editCmd)
}

var editCmd=&cobra.Command{
	Use: "edit",
	Short: "Edit a task",
	Args: cobra.RangeArgs(1,1),
	Run: func(cmd *cobra.Command, args []string){
		id,err := strconv.Atoi(args[0])
		if err!=nil{
			panic(err)
		}
		task,err:=datastore.GetTask(id)
		if err!=nil{
			panic(err)
		}
		if task!=nil{
		fmt.Println(fmt.Sprintf("Current Title : %s",task.Title))
		tempTitle:=""
		fmt.Println("Enter Title (Used Current If Left Empty) : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan(){
			tempTitle=scanner.Text()
		}
		if tempTitle!=""{
			task.Title=tempTitle
		}
		fmt.Println(fmt.Sprintf("Current Description : %s",task.Description))
		fmt.Println("Enter Description ([Default] Empty) : ")
		if scanner.Scan(){
			task.Description=scanner.Text()
		}
		fmt.Println("Status should a [String] in List (Not Started,In Progress,Completed) or Short Hands (ns,ip,c) : ")
		if scanner.Scan(){
			task.Status=scanner.Text()
		}
		switch task.Status{
		case "Not Started","ns":
			task.Status="Not Started"
		case "In Progress","ip":
			task.Status="In Progress"
		case "Completed","c":
			task.Status="Completed"
		default:
			task.Status="Not Started"
		}
		if err:=datastore.EditTasks(id,*task);err!=nil{
			panic(err)
		}
	}
	},
}