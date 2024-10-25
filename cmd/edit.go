package cmd

import (
	"github.com/spf13/cobra"
	datastore "todo/pkg/datastore"
	model "todo/pkg/model"
)

func init(){
	rootCmd.AddCommand(editCmd)
}

var editCmd=&cobra.Command{
	Use: "edit",
	Short: "Edit a task",
}