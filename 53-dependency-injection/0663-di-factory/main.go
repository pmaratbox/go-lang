package main

import (
	"fmt"

	"go.uber.org/dig"
)

// Service is the type produced by the factory. Its value() returns `built`.
type Service struct {
	v string
}

func (s *Service) Value() string { return s.v }

// NewServiceFactory is a FACTORY provider: instead of dig autowiring the
// struct's fields, this function explicitly constructs and configures the
// Service. dig calls it on demand to build the object.
func NewServiceFactory() *Service {
	return &Service{v: "built"}
}

func main() {
	c := dig.New()

	// Register the service via the factory function.
	if err := c.Provide(NewServiceFactory); err != nil {
		panic(err)
	}

	// Resolve the factory-built service and call its method.
	if err := c.Invoke(func(s *Service) {
		fmt.Println(s.Value())
	}); err != nil {
		panic(err)
	}
}
