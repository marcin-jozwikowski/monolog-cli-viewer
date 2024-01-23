package viewer

import (
	"bytes"
	"encoding/json"
	"monolog-cli-viewer/src/cli"
	"monolog-cli-viewer/src/colors"
	"strconv"
	"text/template"
	"time"

	"github.com/stretchr/objx"
)

type LogLine struct {
	json objx.Map
	raw  string
}

var lineAppendix string = "\n"

func init() {
	if *cli.RuntimeConfig.NoNewLine {
		lineAppendix = ""
	}
}

func LogLineFromObjx(json objx.Map, rawLine string) *LogLine {
	l := LogLine{
		json: json,
		raw:  rawLine,
	}
	l.addLevelField()           // fills in the _level field
	l.addColorField()           // fills in the _color, _colorReset, _colors[] fields
	l.addDatetimeField()        // fills in the _datetime field
	l.serializeField("context") // filla in the _context and _contextPretty fields
	l.serializeField("extra")   // fills in the _extra and _extraPretty fields

	return &l
}

func (item *LogLine) GetFormattedString(t *template.Template) string {

	if item.json == nil { // if the item has no json - there was a problem decoding it
		if item.raw != "" { // if raw value has been set - return it in unformatted state
			return item.raw + lineAppendix
		}
		return ""
	}

	var tpl bytes.Buffer
	err2 := t.Execute(&tpl, item.json)
	if err2 != nil {
		panic(err2) // there was an error executing the template and it wasn't in the data
	}

	return tpl.String() + lineAppendix
}

func (item *LogLine) addColorField() {
	colorsMap := colors.NameToColorMap()

	item.json.Set("_colors", colorsMap)
	item.json.Set("_color", colors.GetColorForLogLevel(item.json.Get("_level").Str()))
	item.json.Set("_colorR", colorsMap["reset"])
}

func (item *LogLine) addLevelField() {
	// Define a mapping of Go log levels to PHP Monolog levels
	logLevelMapping := map[string]string{
		"100": "DEBUG",
		"200": "INFO",
		"300": "NOTICE",
		"400": "WARNING",
		"500": "ERROR",
		"600": "CRITICAL",
		"700": "ALERT",
		"800": "EMERGENCY",
	}

	levelInput := item.json.Get("level").String() // get level from the log values
	level, exists := logLevelMapping[levelInput]  // map it
	if !exists {
		level = string(bytes.ToUpper([]byte(levelInput))) // use raw value if not in map - just in uppercase
	}

	item.json.Set("_level", level)
}

func (item *LogLine) serializeField(field string) {
	if !item.json.Has(field) {
		return
	}
	rawField := item.json.Get(field)

	jsonStr, err := json.Marshal(rawField.Data()) // one-line JSON
	if err != nil {
		jsonStr = []byte(err.Error())
	}
	item.json.Set("_"+field, string(jsonStr))

	jsonPretty, err2 := json.MarshalIndent(rawField.Data(), "", "  ") // pretty JSON
	if err2 != nil {
		jsonPretty = []byte(err.Error())
	}
	item.json.Set("_"+field+"Pretty", string(jsonPretty))
}

func (item *LogLine) addDatetimeField() {
	if item.json.Has("timestamp") { // I've seen timestamp field instead of datetime
		i, err := strconv.ParseInt(item.json.Get("timestamp").String(), 10, 64)
		if err != nil {
			return
		}
		tm := time.Unix(i/1000, 0)                                   // timetamp includes ms
		item.json.Set("_datetime", tm.Format("2006-01-02 15:04:05")) // Y-m-d H:i:s

		return
	}

	if item.json.Has("datetime") { // proper monolog time
		date, err2 := time.Parse("2006-01-02T15:04:05.999999999-07:00", item.json.Get("datetime").String())
		if err2 != nil {
			return
		}
		item.json.Set("_datetime", date.Format("2006-01-02 15:04:05")) // Y-m-d H:i:s
	}
}
