package viewer

import (
	"github.com/stretchr/objx"
)

func InitLogLine(rawLine string) LogLine {
	j, err := objx.FromJSON(rawLine)
	if err != nil {
		// cannot have it as JSON so lets at least return the raw line back
		return LogLine{
			raw: rawLine,
		}
	}
	j.Set("_rawLog", rawLine) // raw log preset at _rawLog field

	l := LogLine{
		json: j,
	}
	l.addLevelField()           // fills in the _level field
	l.addColorField()           // fills in the _color, _colorReset, _colors[] fields
	l.addDatetimeField()        // fills in the _datetime field
	l.serializeField("context") // filla in the _context and _contextPretty fields
	l.serializeField("extra")   // fills in the _extra and _extraPretty fields

	return l
}
