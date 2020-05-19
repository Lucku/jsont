// +build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
var Default = Build

var modulePath = "github.com/lucku/jsont"
var name = "jsont"

var ldflags = `-s -w -X "main.timestamp=$TIMESTAMP" -X "main.commitHash=$COMMIT_HASH" -X "main.version=$VERSION"`

// Build the application
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	fileName := name
	if runtime.GOOS == "windows" {
		fileName += ".exe"
	}
	return sh.RunWith(envParams(), "go", "build", "-ldflags", ldflags, "-o", name, modulePath)
}

// Install the application
func Install() error {
	fmt.Println("Installing...")
	return sh.RunWith(envParams(), "go", "install", "-ldflags", ldflags, modulePath)
}

// InstallDeps installs all application's dependencies
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	return sh.RunV("go", "get", "-u", "./...")
}

// Test the application
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "./...")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll(name)
}

// Release the application
func Release() (err error) {
	fmt.Println("Releasing...")
	if os.Getenv("TAG") == "" {
		return errors.New("TAG environment variable is required")
	}
	if err := sh.RunV("git", "tag", "-a", "$TAG"); err != nil {
		return err
	}
	if err := sh.RunV("git", "push", "origin", "$TAG"); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			sh.RunV("git", "tag", "--delete", "$TAG")
			sh.RunV("git", "push", "--delete", "origin", "$TAG")
		}
	}()
	return sh.RunV("goreleaser")
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	tag, _ := sh.Output("git", "describe", "--tags")
	return tag
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "HEAD")
	return hash
}

// envVars returns key/values pairs of environment parameters for the build process
func envParams() map[string]string {

	tag := tag()

	if tag == "" {
		tag = "dev"
	}

	timestamp := time.Now().Format(time.RFC3339)

	return map[string]string{
		"MODULE":      modulePath,
		"TIMESTAMP":   timestamp,
		"VERSION":     tag,
		"COMMIT_HASH": hash(),
	}
}
