package cmd

import (
	"github.com/spf13/cobra"
	datastore "todo/pkg/datastore"
	model "todo/pkg/model"
)

func init(){
	rootCmd.AddCommand(listCmd)
}

var listCmd=&cobra.Command{
	Use: "list",
	Short: "List all tasks",
}