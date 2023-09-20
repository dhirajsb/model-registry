package internal

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var (
	modelsStartToken = regexp.MustCompile("^models:.*")
	modelToken       = regexp.MustCompile("^  ([\\p{L}\\d_]+):.*")
	modelsStopToken  = regexp.MustCompile("^[\\p{L}\\d_]+:.*")
)

type scannerState int

const (
	notModels scannerState = iota
	inModels
	inModelMapping
	doneModels
)

func ScanModelMapping(reader io.Reader, writer io.Writer) (*map[string]string, error) {

	result := make(map[string]string)

	s := bufio.NewScanner(reader)
	w := bufio.NewWriter(writer)
	defer w.Flush()

	state := notModels

	var modelName string
	var modelBuilder strings.Builder

	s.Split(bufio.ScanLines)
	for s.Scan() {

		text := s.Text()

		switch state {
		case notModels:

			if modelsStartToken.MatchString(text) {
				state = inModels
			}
			w.WriteString(text + "\n")

		case inModels:

			submatch := modelToken.FindStringSubmatch(text)
			if submatch != nil {
				state = inModelMapping
				modelName = submatch[1]
				// reset modelBuilder to collect model mapping
				modelBuilder.Reset()
				modelBuilder.WriteString(text + "\n")
			} else if modelsStopToken.MatchString(text) {
				// no model maps??
				state = doneModels
				break
			}

		case inModelMapping:

			// start of another model map?
			submatch := modelToken.FindStringSubmatch(text)
			if submatch != nil {
				// add previous model to map
				result[modelName] = modelBuilder.String()

				// start collecting the new model
				modelName = submatch[1]
				// reset modelBuilder to collect new model mapping
				modelBuilder.Reset()
				modelBuilder.WriteString(text + "\n")

			// start of another top level gqlgen token?
			} else if modelsStopToken.MatchString(text) {
				// add previous model to map
				result[modelName] = modelBuilder.String()
				modelBuilder.Reset()

				state = doneModels

			// non-empty model line
			} else if text != "" {
				// continue to collect model text
				modelBuilder.WriteString(text + "\n")
			}
		}

		if state == doneModels {
			break
		}
	}

	// ran out of lines and still inModelMapping?
	if state == inModelMapping {
		// add last model to map
		result[modelName] = modelBuilder.String()
	}

	return &result, nil
}
