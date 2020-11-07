package gohelper

import (
	"encoding/csv"
	"os"
)

// ReadCSV returns [][]string for
func ReadCSV(filename string, sept rune) ([][]string, error) {
	r := make([][]string, 0)
	f, err := os.Open(filename)
	if err != nil {
		return r, err
	}
	defer f.Close()

	fr := csv.NewReader(f)
	fr.Comma = sept
	return fr.ReadAll()
}
