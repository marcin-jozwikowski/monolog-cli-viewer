package colors

var levelToColorMap = map[string]string{
	"DEBUG":     "white",
	"INFO":      "cyan",
	"NOTICE":    "green",
	"WARNING":   "yellow",
	"ERROR":     "purple",
	"CRITICAL":  "red",
	"ALERT":     "red",
	"EMERGENCY": "red",
}

var nameToColorMap = map[string]string{
	"reset":     "\033[0m",
	"bold":      "\033[1m",
	"underline": "\033[4m",
	"strike":    "\033[9m",
	"italic":    "\033[3m",
	"red":       "\033[31m",
	"green":     "\033[32m",
	"yellow":    "\033[33m",
	"blue":      "\033[34m",
	"purple":    "\033[35m",
	"cyan":      "\033[36m",
	"white":     "\033[37m",
}

var nameToColorMapEmpty = map[string]string{
	"reset":     "",
	"bold":      "",
	"underline": "",
	"strike":    "",
	"italic":    "",
	"red":       "",
	"green":     "",
	"yellow":    "",
	"blue":      "",
	"purple":    "",
	"cyan":      "",
	"white":     "",
}

var colorsEnabled bool

func SetEnabled(enabled bool) {
	colorsEnabled = enabled
}

// the full map of all colors
func NameToColorMap() map[string]string {
	if colorsEnabled == true {
		return nameToColorMap
	}

	// functionally disable the colors while leaving templates running
	return nameToColorMapEmpty
}

func GetColorForLogLevel(name string) string {
	levelColor, exitsts := levelToColorMap[name]
	if !exitsts {
		// in case there's an unaccounted for level
		levelColor = "white"
	}

	return NameToColorMap()[levelColor]
}
