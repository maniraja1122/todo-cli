package cmd

import (
	"github.com/spf13/cobra"
	datastore "todo/pkg/datastore"
	model "todo/pkg/model"
)

func init(){
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd=&cobra.Command{
	Use: "delete",
	Short: "Delete a task",
}