package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucku/jsont/cmd"
	"github.com/lucku/jsont/transform"
	"github.com/tidwall/gjson"
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

			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {

	bigFileData, _ := ioutil.ReadFile("reference-data/big.json")

	parsed := gjson.ParseBytes(bigFileData)

	j := transform.JSONIterator{Data: &parsed}

	for j.Next() {
		fmt.Println(j.Value().Path, j.Value().Value)
	}
}
