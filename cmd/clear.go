package cmd

import (
	datastore "todo/pkg/datastore"

	"github.com/spf13/cobra"
)

func init(){
	clearCmd.Flags().StringP("status","s","","Only delete tasks with this status")
	rootCmd.AddCommand(clearCmd)
}

var clearCmd=&cobra.Command{
	Use: "clear",
	Short: "Clear All Tasks",
	Long: "Clear will remove all the tasks",
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string){
		flagStatus:=cmd.Flag("status")
		if flagStatus.Value.String()!=""{
			delStatus:=""
			switch flagStatus.Value.String(){
			case "Not Started","ns":
				delStatus="Not Started"
			case "In Progress","ip":
				delStatus="In Progress"
			case "Completed","c":
				delStatus="Completed"
			}
			if delStatus==""{
				println(`Please Enter a Proper Status - [String] in List ("Not Started","In Progress","Completed") or Short Hands (ns,ip,c)`)
			} else{
				if err:=datastore.DeleteTasksByStatus(delStatus); err!=nil{
					panic(err)
				}
			}
		} else{
			if err:=datastore.DeleteAllTasks();err!=nil{
				panic(err)
			}
		}
	},
}