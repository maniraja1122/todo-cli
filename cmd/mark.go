package cmd

import (
	"strconv"
	datastore "todo/pkg/datastore"

	"github.com/spf13/cobra"
)

func init(){
	markCmd.Flags().StringP("status","s","Completed","Status of the task")
	markCmd.MarkFlagRequired("status")
	rootCmd.AddCommand(markCmd)
}

var markCmd=&cobra.Command{
	Use: "mark",
	Short: "Mark the status of a task",
	Long: `
	Mark the status of a task
	Status can be a String in List ("Not Started","In Progress","Completed") - Make sure to Add the Double Quotes ""
	or you can use Short Hands for each possible status value (ns,ip,c)
	By Default, It will be marked as Completed
	`,
	Args: cobra.RangeArgs(1,1),
	Run: func(cmd *cobra.Command,args []string){
		id,err := strconv.Atoi(args[0])
		if err!=nil{
			panic(err)
		}
		markStatus:=""
		statusFlag:=cmd.Flag("status")
		switch statusFlag.Value.String(){
		case "Not Started","ns":
			markStatus="Not Started"
		case "In Progress","ip":
			markStatus="In Progress"
		case "Completed","c":
			markStatus="Completed"
		default:
			markStatus="Completed"
		}
		if err:=datastore.MarkTask(id,markStatus); err!=nil{
			panic(err)
		}
		},
}