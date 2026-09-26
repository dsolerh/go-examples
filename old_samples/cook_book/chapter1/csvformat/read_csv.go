package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie will hold the parsed csv
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV gives some examples of processing csv
// that is passed in as an io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)
	// these are some optional configuration options
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// grab and ignore the header for now
	// this can be used for dictionary keys or some other forms of lookup
	_, err := r.Read()
	if err != nil {
		return nil, err
	}

	// loop until it's all processed
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}
		movies = append(movies, Movie{record[0], record[1], int(year)})
	}

	return movies, nil
}

// AddMoviesFromText uses the csv parser with a string
func AddMoviesFromText() error {
	// this is an example of us taking a string, converting
	// it to a buffer, and reading it
	// with the package
	in := `
- First our headers
movie title;director;year released
- then the data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", m)
	return nil
}
