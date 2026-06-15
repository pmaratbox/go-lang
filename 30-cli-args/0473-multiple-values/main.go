package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var nums []int
	cmd := &cobra.Command{Use: "app", Run: func(cmd *cobra.Command, args []string) {
		sum := 0
		for _, n := range nums {
			sum += n
		}
		fmt.Println(sum)
	}}
	cmd.Flags().IntSliceVar(&nums, "num", nil, "")
	cmd.SetArgs([]string{"--num", "1", "--num", "2", "--num", "3"})
	cmd.Execute()
}
