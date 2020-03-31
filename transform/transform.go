package transform

import "fmt"

func Transform(inFile string, outFile string, transFile string) error {

	fmt.Printf("Transform %s with %s to %s", inFile, outFile, transFile)

	return nil
}

func loadInputJSONFile() {

}
