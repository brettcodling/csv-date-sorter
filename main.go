package main

import (
	"encoding/csv"
	"flag"
	"os"
	"sort"
	"time"
)

func main() {
	filename := flag.String("filename", "", "The filename")
	column := flag.Int("column", 0, "The column number to sort by")
	hasHeader := flag.Bool("hasheader", false, "Whether the file has a header")
	output := flag.String("output", "sorted.csv", "The output file")
	reverse := flag.Bool("reverse", false, "Whether to sort the file in reverse")
	dateFormat := flag.String("format", "2006-01-02 15:04:05", "The date format to use. Note this is the golang magic date")
	flag.Parse()

	infile, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	outfile, err := os.Create(*output)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(infile)
	writer := csv.NewWriter(outfile)
	var lineNo int
	var data [][]string
	for {
		row, err := reader.Read()
		if err != nil {
			_, ok := err.(*csv.ParseError)
			if ok {
				panic("Failed to parse input file")
			}
			break
		}
		if lineNo == 0 && *hasHeader {
			writer.Write(row)
		} else {
			data = append(data, row)
		}
		lineNo++
	}

	sort.Slice(data, func(i, j int) bool {
		iTime, _ := time.Parse(*dateFormat, data[i][*column])
		jTime, _ := time.Parse(*dateFormat, data[j][*column])
		if *reverse {
			return iTime.After(jTime)
		}
		return iTime.Before(jTime)
	})

	for _, row := range data {
		writer.Write(row)
	}
	writer.Flush()
	if writer.Error() != nil {
		panic(writer.Error().Error())
	}
}
