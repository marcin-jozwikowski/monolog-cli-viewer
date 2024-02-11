package reader

import (
	"errors"
	"regexp"
	"strings"

	"github.com/stretchr/objx"
)

var monologRegex *regexp.Regexp
var jsonSegmentsRegex []*regexp.Regexp

func init() {
	monologRegex = regexp.MustCompile(`\[(?P<time>[\S]+)\] (?P<channel>[\S]+)\.(?P<level>[\S]+):`)
	jsonSegmentsRegex = []*regexp.Regexp{
		regexp.MustCompile(`\{(?:[^{}]*\{[^{}]*\}[^{}]*)*\}|\[.*?\]`),
		regexp.MustCompile(`\{.*?\}|\[.*?\]`),
		regexp.MustCompile(`(\{[^}]*\}|\[[^]]*\])`),
	}
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

	contextExtra, message := extractJSONSegments(messageContextExtra)
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

func extractJSONSegments(input string) ([]string, string) {
	// @todo further improve on the algorithm to consider only one type array/object at once
	result := []string{}
	inputLength := len(input)
	openenedCount := 0
	closingIndex := inputLength
	for index := inputLength - 1; index > 0; index-- {
		currentChar := input[index]
		if currentChar == '}' || currentChar == ']' { // we're at the closing char
			if openenedCount == 0 { // we're not inside JSON
				closingIndex = index // means we've found the outermost closing tag
			}
			openenedCount++
		}

		if currentChar == '{' || currentChar == '[' { // we're at the opening char
			openenedCount--
			if openenedCount == 0 {
				// prepend the JSON to the result. We're reading backwards.
				result = append([]string{input[index : closingIndex+1]}, result...)
				input = input[0:index]

				if len(result) == 2 { // we got context, and extra - enough
					break
				}
			}
		}
	}

	return result, strings.Trim(input, "\r\t\n ")
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
