package names

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
)

//go:embed presets/common/*.csv
var commonPresets embed.FS

// Common holds a list of common names and their frequency.
type Common struct {
	names []string
	freqs []float32
}

// NewCommonPreset creates a new list of common names based on the provided
// preset.
//
// The available presets are provided in the presets/common folder. The name
// must be provided without the .csv suffix, e.g. "US_FIRST_NAME".
//
// Providing an invalid preset name will result in an error.
func NewCommonPreset(preset string) (*Common, error) {
	f, err := commonPresets.Open(filepath.Clean(fmt.Sprintf("presets/common/%v.csv", preset)))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	return NewCommon(f)
}

// NewCommon create a new list of common names based on the provided reader.
//
// The reader must contain CSV data with exactly two columns and no header. The
// first column must contain the name and the second value must contain a 32-bit
// float with the frequency (likeliness of that name). Furthermore, the data
// must already be in the correct order with the highest frequency at the start.
//
// If the reader returns an error or the format of the data is not correct, an
// error is returned.
func NewCommon(r io.Reader) (*Common, error) {
	names := []string{}
	freqs := []float32{}
	cr := csv.NewReader(r)
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(row) != 2 {
			return nil, fmt.Errorf("invalid number of columns: %d", len(row))
		}
		freq, err := strconv.ParseFloat(row[1], 32)
		if err != nil {
			return nil, err
		}

		names = append(names, row[0])
		freqs = append(freqs, float32(freq))
	}

	return &Common{
		names: names,
		freqs: freqs,
	}, nil
}

// Top returns the n most frequent names.
//
// If n is bigger than the list of names, then the whole list will be returned.
// For values of n less than 1 and empty list will be returned.
func (t *Common) Top(n int) []string {
	return t.names[:max(0, min(n, len(t.names)))]
}

// Frequency returns the frequency for the provided name.
//
// If the name was not found in the list, then a frequency of 0.0 will be returned.
func (t *Common) Frequency(name string) float32 {
	for i := range t.names {
		if t.names[i] == name {
			return t.freqs[i]
		}
	}
	return 0
}
