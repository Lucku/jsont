package main

import (
	"fmt"
)

var (
	timestamp  string
	version    string
	commitHash string
)

// VersionCmd outputs the version, timestamp and commit hash of the jsont build
type VersionCmd struct {
}

// Name returns the name of the version subcommand
func (v *VersionCmd) Name() string {
	return "version"
}

// Init initializes the version subcommand
func (v *VersionCmd) Init(args []string) error {
	return nil
}

// Run executes the version subcommand
func (v *VersionCmd) Run() error {
	fmt.Printf("jsont version %s %s (build date: %s)\n", version, commitHash, timestamp)
	return nil
}

// PrintUsage returns a description of the version command's usage
func (v *VersionCmd) PrintUsage() {
}

// NewVersionCmd returns an instance of a version subcommand
func NewVersionCmd() *VersionCmd {
	return new(VersionCmd)
}
