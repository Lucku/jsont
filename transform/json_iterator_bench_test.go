package transform

import (
	"io/ioutil"
	"testing"

	"github.com/tidwall/gjson"
)

func readJSONFile(file string) []byte {
	bytes, _ := ioutil.ReadFile(file)
	return bytes
}

func BenchmarkIterateJSONFile(b *testing.B) {

	filepath := "../reference-data/"

	files := []string{"medium.json", "big.json"}

	for _, file := range files {
		b.Run(file, func(b *testing.B) {
			benchmarkIterateJSONFile(b, filepath+file)
		})
	}
}

func benchmarkIterateJSONFile(b *testing.B, file string) {
	b.StopTimer()
	bytes := readJSONFile(file)

	data := gjson.ParseBytes(bytes)

	j := JSONIterator{Data: &data}

	b.StartTimer()
	b.ReportAllocs()
	b.SetBytes(int64(len(bytes)))

	for j.Next() {
	}

}
