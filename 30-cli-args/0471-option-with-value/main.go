package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var name string
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(name)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "a name to print")

	// Parse a HARDCODED argv (not os.Args) so output is deterministic.
	cmd.SetArgs([]string{"--name", "alice"})
	cmd.Execute()
}
