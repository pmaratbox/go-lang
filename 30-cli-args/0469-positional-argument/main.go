package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:  "app name",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Println(name)
		},
	}
	cmd.SetArgs([]string{"alice"})
	cmd.Execute()
}
