package cmd

import (
	"errors"
	"flag"
)

type TransformCmd struct {
	fs        *flag.FlagSet
	inFile    string
	outFile   string
	transFile string
}

func (t *TransformCmd) Name() string {
	return t.fs.Name()
}

func (t *TransformCmd) Init(args []string) error {

	if err := t.fs.Parse(args); err != nil {
		return err
	}

	if t.outFile == "" {
		return errors.New("No output file provided")
	}

	if t.transFile == "" {
		return errors.New("No transformation file provided")
	}

	if t.fs.NArg() < 1 {
		return errors.New("No input file provided")
	}

	t.inFile = t.fs.Args()[1]

	return nil
}

func (t *TransformCmd) Run() error {
	return nil
}

func (t *TransformCmd) PrintUsage() {
	t.fs.PrintDefaults()
}

func NewTransformCmd() *TransformCmd {

	tc := &TransformCmd{
		fs: flag.NewFlagSet("transform", flag.ContinueOnError),
	}

	tc.fs.StringVar(&tc.outFile, "o", "", "The file where output is written to (required)")
	tc.fs.StringVar(&tc.transFile, "t", "", "Transformation file (required)")

	return tc
}
