package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use : "todo",
	Aliases: []string{"td"},
	Short : "Your CLI Todo Tracker",
	Long: "Use this tool to add, list, manage and track your day-to-day priority tasks without leaving the CLI.",
	Run: func (cmd *cobra.Command,arg []string){
	},
}

func Execute(){
	if err:=rootCmd.Execute(); err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}