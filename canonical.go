package names

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

//go:embed presets/canonical/*.txt
var canonicalPresets embed.FS

// Canonical holds a map of names and their canonical representation
type Canonical struct {
	values map[string]string
}

// NewCanonicalPreset creates a new canonical map based on the provided preset.
//
// The available presets are provided in the presets/canonical folder. The name
// must be provided without the .txt suffix, e.g. "NICKNAME".
//
// Providing an invalid preset name will result in an error.
func NewCanonicalPreset(preset string) (*Canonical, error) {
	f, err := canonicalPresets.Open(filepath.Clean(fmt.Sprintf("presets/canonical/%v.txt", preset)))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	return NewCanonical(f)
}

// NewCanonical creates a new canonical map based on the provided reader.
//
// The reader must contain lines with at least two comma separated names. The
// first name is considered the canonical name and all further names are aliases
// that shall be matched to that name.
//
// It depends on the data to define what a "canonical name" means.
//
// If the reader returns an error or the format of the data is not correct, an
// error is returned.
func NewCanonical(r io.Reader) (*Canonical, error) {
	scanner := bufio.NewScanner(r)
	values := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, ",")
		if len(names) < 2 {
			return nil, fmt.Errorf("invalid number of entries: %d (line: %v)", len(names), line)
		}
		n, cc := names[0], names[1:]
		for _, c := range cc {
			values[c] = n
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Canonical{
		values: values,
	}, nil
}

// Of returns the canonical name for a given name or the name itself if no
// canonical name exists.
func (c *Canonical) Of(name string) string {
	if name, ok := c.values[name]; ok {
		return name
	}
	return name
}
