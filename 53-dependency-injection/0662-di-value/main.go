package main

import (
	"fmt"

	"go.uber.org/dig"
)

// Cfg is a constant value (configuration) we register with the container.
type Cfg struct{ V string }

// Service receives the injected Cfg and returns its value.
type Service struct{ cfg Cfg }

// NewCfg provides the constant value v1.
func NewCfg() Cfg { return Cfg{V: "v1"} }

// NewService is the constructor dig uses; it depends on Cfg.
func NewService(c Cfg) *Service { return &Service{cfg: c} }

// Value returns the injected configuration value.
func (s *Service) Value() string { return s.cfg.V }

func main() {
	c := dig.New()

	// Register the constant value and the service that consumes it.
	if err := c.Provide(NewCfg); err != nil {
		panic(err)
	}
	if err := c.Provide(NewService); err != nil {
		panic(err)
	}

	// Resolve the service from the container and print the injected value.
	if err := c.Invoke(func(s *Service) {
		fmt.Println(s.Value())
	}); err != nil {
		panic(err)
	}
}
