package cmd

import (
	"strconv"
	datastore "todo/pkg/datastore"

	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(showCmd)
}

var showCmd=&cobra.Command{
	Use: "show",
	Short: "Show a task",
	Args: cobra.RangeArgs(1,1),
	Run: func(cmd *cobra.Command, args []string){
		id,err := strconv.Atoi(args[0])
		if err!=nil{
			panic(err)
		}
		if err:=datastore.ShowTask(id);err!=nil{
			panic(err)
		}
	},
}