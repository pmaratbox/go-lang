package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var verbose bool
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(verbose)
		},
	}
	cmd.Flags().BoolVar(&verbose, "verbose", false, "verbose output")
	cmd.SetArgs([]string{"--verbose"})
	cmd.Execute()
}
