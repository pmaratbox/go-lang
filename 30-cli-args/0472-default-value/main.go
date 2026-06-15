package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var count int
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(count)
		},
	}
	cmd.Flags().IntVar(&count, "count", 1, "")
	// Parse a hardcoded EMPTY argv for deterministic output.
	cmd.SetArgs([]string{})
	cmd.Execute()
}
