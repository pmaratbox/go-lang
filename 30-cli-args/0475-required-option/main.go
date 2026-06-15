package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var id int
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(id)
		},
	}
	cmd.Flags().IntVar(&id, "id", 0, "")
	cmd.MarkFlagRequired("id")
	cmd.SetArgs([]string{"--id", "42"})
	cmd.Execute()
}
