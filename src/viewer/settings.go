package viewer

type Settings struct {
	NoNewLine           bool
	ShowFileChangeLine  bool
	ShowParsedLinesOnly bool
}

var settings Settings = Settings{
	NoNewLine:           false,
	ShowFileChangeLine:  false,
	ShowParsedLinesOnly: false,
}

func SetSettings(sets Settings) {
	settings = sets
}
