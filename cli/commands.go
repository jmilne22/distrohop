package cli

import "fmt"

type Globals struct {
	Config string `help:"config file to use" default:"./config/config.yaml"`
}

type StateCmd struct{}

func (e *StateCmd) Run(g *Globals) error {
	fmt.Println("Hello world")
	return nil
}
