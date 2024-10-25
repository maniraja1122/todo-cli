package cmd

import (
	"github.com/spf13/cobra"
	datastore "todo/pkg/datastore"
	model "todo/pkg/model"
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
	Status can be a String in List ("Not Started","In Progress","Completed")
	or you can use Short Hands for each possible status value ("ns","ip","c")
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command,args []string){
		for _,arg := range args{

		}
	},
}