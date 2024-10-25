package cmd

import (
	"github.com/spf13/cobra"
	datastore "todo/pkg/datastore"
	model "todo/pkg/model"
)

func init(){
	rootCmd.AddCommand(showCmd)
}

var showCmd=&cobra.Command{
	Use: "show",
	Short: "Show a task",
}