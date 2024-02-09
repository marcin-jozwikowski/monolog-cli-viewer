package viewer

import (
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
	j, err := objx.FromJSON(rawLine)
	if err != nil {
		j, err = reader.MonologFormat(rawLine)
		if err != nil {
			// cannot have it as JSON so lets at least return the raw line back
			return &LogLine{
				raw: strings.Trim(rawLine, " \r\n\t"),
			}
		}
	}

	return LogLineFromObjx(j, rawLine)
}

func isFileChangeLine(rawLine string) bool {
	found := fileChangeRegex.FindString(rawLine)

	return len(found) > 0
}
