package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lucku/jsont/cmd"
)

type runner interface {
	Init([]string) error
	Run() error
	Name() string
	PrintUsage()
}

func execute(args []string) error {

	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []runner{
		cmd.NewTransformCmd(),
		cmd.NewVersionCmd(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {

			if err := cmd.Init(os.Args[2:]); err != nil {
				cmd.PrintUsage()
				return err
			}

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("Error on execution of subcommand %s: %v", subcommand, err)
			}

			return nil
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {

	if err := execute(os.Args[1:]); err != nil {
		fmt.Printf("%v\n", err)
	}
}
