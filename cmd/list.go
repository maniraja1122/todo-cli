package cmd

import (
	datastore "todo/pkg/datastore"

	"github.com/spf13/cobra"
)

func init(){
	listCmd.Flags().StringP("status","s","","Only List Tasks with this status")
	rootCmd.AddCommand(listCmd)
}

var listCmd=&cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string){
       statusFlag:=cmd.Flag("status")
	   if statusFlag.Value.String()!=""{
		searchStatus:=""
		switch statusFlag.Value.String(){
		case "Not Started","ns":
			searchStatus="Not Started"
		case "In Progress","ip":
			searchStatus="In Progress"
		case "Completed","c":
			searchStatus="Completed"
		default:
			searchStatus="Not Started"
		}
		if err:=datastore.ListTaskbyStatus(searchStatus);err!=nil{
			panic(err)
		}
	   } else{
		if err:=datastore.ListAllTasks();err!=nil{
			panic(err)
		}
	   }
	},

}