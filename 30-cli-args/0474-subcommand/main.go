package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	addCmd := &cobra.Command{
		Use:  "add",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			a, _ := strconv.Atoi(args[0])
			b, _ := strconv.Atoi(args[1])
			fmt.Println(a + b)
		},
	}

	root := &cobra.Command{Use: "app"}
	root.AddCommand(addCmd)

	// Hardcoded argv for deterministic output (not os.Args).
	root.SetArgs([]string{"add", "2", "3"})
	root.Execute()
}
