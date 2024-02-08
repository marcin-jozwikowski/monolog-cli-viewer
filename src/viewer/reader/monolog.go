package reader

import (
	"errors"
	"regexp"
	"strings"

	"github.com/stretchr/objx"
)

var monologRegex *regexp.Regexp

func init() {
	monologRegex = regexp.MustCompile(`\[(?P<time>[\S]+)\] (?P<channel>[\S]+)\.(?P<level>[\S]+):`)
}

func MonologFormat(rawLine string) (objx.Map, error) {
	mappedValues, matched := regexMatchToMap(rawLine, monologRegex) // extract time, channel, and level from raw string

	if len(mappedValues) == 0 {
		// nothing was found - not a Monolog format
		return objx.MSI(), errors.New("could not read")
	}

	messageContextExtra, removed := strings.CutPrefix(rawLine, matched)
	if !removed {
		// does not have
		return objx.MSI(), errors.New("no context nor extra")
	}

	contextExtra := extractJSONSegments(messageContextExtra)
	message, _ := strings.CutSuffix(messageContextExtra, strings.Join(contextExtra, " "))
	message = strings.Trim(message, " \r\n\t")
	objxFields := map[string]objx.Map{}

	if len(contextExtra) > 0 && len(contextExtra[0]) > 2 { // at least one field - context
		objxFields["context"] = getJsonOrEmpty(contextExtra[0])
	}

	if len(contextExtra) > 1 && len(contextExtra[1]) > 2 { // more than one field - this mean extra
		objxFields["extra"] = getJsonOrEmpty(contextExtra[1])
	}

	return objx.MSI(
		"level_name", mappedValues["level"],
		"level", mappedValues["level"],
		"datetime", mappedValues["time"],
		"message", message,
		"channel", mappedValues["channel"],
		"context", objxFields["context"],
		"extra", objxFields["extra"],
	), nil
}

func getJsonOrEmpty(input string) objx.Map {
	if obj, err := objx.FromJSON(input); err == nil {
		return obj
	}
	return objx.MSI()
}

func extractJSONSegments(input string) []string {
	// Define a regular expression to match JSON objects, including nested structures
	jsonRegex := `\{(?:[^{}]*\{[^{}]*\}[^{}]*)*\}|\[.*?\]`

	// Compile the regular expression
	re := regexp.MustCompile(jsonRegex)

	// Find all matches in the input string
	matches := re.FindAllString(input, -1)

	return matches
}

func regexMatchToMap(value string, regex *regexp.Regexp) (map[string]string, string) {
	match := regex.FindString(value)
	matches := regex.FindStringSubmatch(value)
	results := make(map[string]string)

	if len(match) > 0 {
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				results[name] = matches[i]
			}
		}
	}

	return results, match
}
