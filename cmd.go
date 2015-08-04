package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var Cmd = &cobra.Command{
		Use:   "cmd",
		Short: "cmd a short description here",
		Long: `A very long description written here
and included here
and here`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}
	Cmd.AddCommand(versionCmd)
	Cmd.Execute()
}
