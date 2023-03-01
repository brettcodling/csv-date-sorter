package sort

import (
	"encoding/csv"
	"os"
	"sort"
	"time"
)

// Sort will sort the given file using the passed parameters
func Sort(input string, output string, column int, dateFormat string, hasHeader bool, reverse bool) {
	infile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	outfile, err := os.Create(output + "temp")
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
		if lineNo == 0 && hasHeader {
			writer.Write(row)
		} else {
			data = append(data, row)
		}
		lineNo++
	}

	sort.Slice(data, func(i, j int) bool {
		iTime, _ := time.Parse(dateFormat, data[i][column])
		jTime, _ := time.Parse(dateFormat, data[j][column])
		if reverse {
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

	os.Rename(output+"temp", output)
}
