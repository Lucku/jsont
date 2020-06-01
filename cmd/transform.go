package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/lucku/jsont"
)

// TransformCmd transforms an input JSON using a transformation description
type TransformCmd struct {
	fs        *flag.FlagSet
	inFiles   []string
	outDir    string
	transFile string
}

// Name returns the name of the transform subcommand
func (t *TransformCmd) Name() string {
	return t.fs.Name()
}

// Init initializes the transform subcommand
func (t *TransformCmd) Init(args []string) error {

	if err := t.fs.Parse(args); err != nil {
		return err
	}

	if t.transFile == "" {
		return errors.New("No transformation file provided")
	}

	if t.fs.NArg() < 1 {
		return errors.New("No input file(s) provided")
	}

	t.inFiles = t.fs.Args()

	return nil
}

// Run executes the transform subcommand
func (t *TransformCmd) Run() error {

	jt, err := jsont.NewJSONTransformerWithFile(t.transFile)

	if err != nil {
		return err
	}

	for _, f := range t.inFiles {

		outData, err := jt.TransformWithFile(f)

		if err != nil {
			return err
		}

		fileName := filepath.Base(f)
		dotTokens := strings.Split(fileName, ".")
		outFileName := strings.Join(dotTokens[:len(dotTokens)-1], ".") + ".out." + dotTokens[len(dotTokens)-1]

		outFile := filepath.Join(t.outDir, outFileName)

		if err := ioutil.WriteFile(outFile, outData, 0644); err != nil {
			return err
		}
	}

	return nil
}

// PrintUsage returns a description of the transform command's usage
func (t *TransformCmd) PrintUsage() {
	t.fs.PrintDefaults()
}

// NewTransformCmd returns an instance of a transform subcommand
func NewTransformCmd() *TransformCmd {

	tc := &TransformCmd{
		fs: flag.NewFlagSet("transform", flag.ContinueOnError),
	}

	tc.fs.StringVar(&tc.outDir, "o", "", "The directory where output file are written to")
	tc.fs.StringVar(&tc.transFile, "t", "", "Transformation file (required)")

	return tc
}
