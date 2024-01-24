package viewer

import "text/template"

type Settings struct {
	NoNewLine           bool
	ShowFileChangeLine  bool
	ShowParsedLinesOnly bool
	Template            *template.Template
}

var settings Settings = Settings{
	NoNewLine:           false,
	ShowFileChangeLine:  false,
	ShowParsedLinesOnly: false,
}

func SetSettings(sets Settings) {
	settings = sets
}
