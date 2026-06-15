package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// choiceValue implements pflag.Value to restrict a flag to a fixed set of
// choices, giving cobra a typed "enum" option.
type choiceValue struct {
	value   string
	allowed []string
}

func (c *choiceValue) String() string { return c.value }

func (c *choiceValue) Type() string { return "choice" }

func (c *choiceValue) Set(v string) error {
	for _, a := range c.allowed {
		if v == a {
			c.value = v
			return nil
		}
	}
	return fmt.Errorf("must be one of {%s}", strings.Join(c.allowed, ", "))
}

func main() {
	color := &choiceValue{allowed: []string{"red", "green", "blue"}}
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(color.value)
		},
	}
	var _ pflag.Value = color
	cmd.Flags().Var(color, "color", "one of {red, green, blue}")

	// Parse a HARDCODED argv (not os.Args) so output is deterministic.
	cmd.SetArgs([]string{"--color", "green"})
	cmd.Execute()
}
