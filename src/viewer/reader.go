package viewer

import (
	"errors"
	"regexp"
	"strings"

	"github.com/stretchr/objx"
)

var monologRegex *regexp.Regexp
var fileChangeRegex *regexp.Regexp

func init() {
	monologRegex = regexp.MustCompile(`\[(?P<time>[\S]+)\] (?P<channel>[\S]+)\.(?P<level>[\S]+): (?P<message>[\w\d\s]+)`)
	fileChangeRegex = regexp.MustCompile(`^==>.*<==$`)
}

func InitLogLine(rawLine string) *LogLine {
	j, err := objx.FromJSON(rawLine)
	if err != nil {
		j, err = readMonologFormat(rawLine)
		if err != nil {
			// cannot have it as JSON so lets at least return the raw line back
			return &LogLine{
				raw: strings.Trim(rawLine, " \r\n\t"),
			}
		}
	}

	return LogLineFromObjx(j, rawLine)
}

func readMonologFormat(rawLine string) (objx.Map, error) {
	results, matched := regexMatchToMap(rawLine, monologRegex)

	if len(results) == 0 {
		return objx.MSI(), errors.New("could not read")
	}

	contextExtra, removed := strings.CutPrefix(rawLine, matched)
	if !removed {
		return objx.MSI(), errors.New("no context nor extra")
	}

	contextExtra = strings.ReplaceAll(contextExtra, "\r\n", " ")     // remove Windows-style newline
	contextExtra = strings.ReplaceAll(contextExtra, "}\n", "} ")     // remove Unix newline
	contextExtra = strings.ReplaceAll(contextExtra, "} []", "}\n{}") // empty extra
	contextExtra = strings.ReplaceAll(contextExtra, "} {", "}\n{")   // filled extra

	splitContextExtra := strings.Split(contextExtra, "\n")
	objxFields := map[string]objx.Map{}

	if len(splitContextExtra) > 0 { // at least one field - context
		objxFields["context"] = objx.MustFromJSON(splitContextExtra[0])
	} else { // no additional fields
		objxFields["context"] = objx.MSI()
	}

	if len(splitContextExtra) > 1 { // more than one field - this mean extra
		objxFields["extra"] = objx.MustFromJSON(splitContextExtra[1])
	} else { // no extras
		objxFields["extra"] = objx.MSI()
	}

	return objx.MSI(
		"level_name", results["level"],
		"level", results["level"],
		"datetime", results["time"],
		"message", results["message"],
		"channel", results["channel"],
		"context", objxFields["context"],
		"extra", objxFields["extra"],
	), nil
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

func isFileChangeLine(rawLine string) bool {
	found := fileChangeRegex.FindString(rawLine)

	return len(found) > 0
}
