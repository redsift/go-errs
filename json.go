package errs

import (
	"encoding/json"
	"fmt"
	"strings"
)

func snippet(data []byte, offset int, message string) string {
	const snippetRange = 32
	const lb = "\n"
	s := offset - snippetRange
	e := offset + snippetRange
	if s < 0 {
		s = 0
	}

	if e > len(data) {
		e = len(data)
	}

	latter := string(data[offset:e])
	latterLb := strings.Index(latter, lb)
	if latterLb != -1 {
		e = offset + latterLb
	}

	first := string(data[:offset])
	firstLb := strings.LastIndex(first, lb)

	if firstLb != -1 {
		s = firstLb + 1
	}

	ln := fmt.Sprintf("line:%02d", strings.Count(first, lb)+1)

	line := string(data[s:e])
	padding := strings.Repeat(" ", (offset-s)+len(ln))

	return fmt.Sprintf("\n%s %s\n%s ^ %s", ln, line, padding, message)
}

// Take a JSON error and try and improve the presentation for console
func PostProcessJsonError(data []byte, err error) error {
	switch v := err.(type) {
	case *json.SyntaxError:
		snip := snippet(data, int(v.Offset), v.Error())
		return fmt.Errorf("Could not parse JSON. %s", snip)
	case *json.UnmarshalTypeError:
		snip := snippet(data, int(v.Offset), fmt.Sprintf("%s is not parseable as an %s", v.Value, v.Type))
		return fmt.Errorf("Could not parse JSON. %s", snip)
	}
	return err
}
