package viewer

import (
	"errors"
	"monolog-cli-viewer/src/viewer/reader"
	"regexp"
	"strings"

	"github.com/stretchr/objx"
)

var fileChangeRegex *regexp.Regexp

func init() {
	fileChangeRegex = regexp.MustCompile(`^==>.*<==$`)
}

func InitLogLine(rawLine string) *LogLine {
	var j objx.Map
	var err error
	rawLine = strings.Trim(rawLine, "\r\n\t ")
	if len(rawLine) == 0 {
		// empty line
		return unparsedLine(rawLine)
	}

	if rawLine[0] == '{' { // if it starts with a JSON opening
		j, err = objx.FromJSON(rawLine) // lets try parsing it as JSON
	} else {
		err = errors.New("not json")
	}

	if err != nil { // JSON did not work or couldn't work
		j, err = reader.MonologFormat(rawLine)
		if err != nil {
			// cannot have it as JSON so lets at least return the raw line back
			return unparsedLine(rawLine)
		}
	}

	return LogLineFromObjx(j, rawLine)
}

func isFileChangeLine(rawLine string) bool {
	found := fileChangeRegex.FindString(rawLine)

	return len(found) > 0
}

func unparsedLine(rawLine string) *LogLine {
	return &LogLine{
		raw: strings.Trim(rawLine, " \r\n\t"),
	}
}
