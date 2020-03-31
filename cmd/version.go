package cmd

import (
	"fmt"
)

const _version = "0.1"

type VersionCmd struct {
}

func (v *VersionCmd) Name() string {
	return "version"
}

func (v *VersionCmd) Init(args []string) error {
	return nil
}

func (v *VersionCmd) Run() error {
	fmt.Printf("jsont version %s\n", _version)
	return nil
}

func (v *VersionCmd) PrintUsage() {
}

func NewVersionCmd() *VersionCmd {
	return new(VersionCmd)
}
