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
	Long: `
  _____ ___  ____   ___
 |_   _/ _ \|  _ \ / _ \
   | || | | | | | | | | |
   | || |_| | |_| | |_| |
   |_| \___/|____/ \___/
   
Use this tool to add, list, manage and track your day-to-day priority tasks without leaving the CLI.

https://github.com/maniraja1122/todo-cli
	`,
}

func Execute(){
	if err:=rootCmd.Execute(); err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}