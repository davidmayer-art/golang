package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var root = cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "app [source] [description] ",
		Short:                 "Permet de copier des fichier",
		Args:                  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			destination := args[1]

			fmt.Println(source, destination)
		},
	}

	root.Execute()
}
