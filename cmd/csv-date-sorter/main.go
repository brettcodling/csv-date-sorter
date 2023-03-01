package main

import (
	"flag"

	"github.com/brettcodling/csv-date-sorter/sort"
)

func main() {
	input := flag.String("input", "", "The input filename")
	output := flag.String("output", "sorted.csv", "The output filename")
	column := flag.Int("column", 0, "The column number to sort by")
	dateFormat := flag.String("format", "2006-01-02 15:04:05", "The date format to use. Note this is the golang magic date")
	hasHeader := flag.Bool("hasheader", false, "Whether the file has a header")
	reverse := flag.Bool("reverse", false, "Whether to sort the file in reverse")
	flag.Parse()

	sort.Sort(*input, *output, *column, *dateFormat, *hasHeader, *reverse)
}
