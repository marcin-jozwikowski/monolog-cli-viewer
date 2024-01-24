package cli

import (
	"flag"
	"monolog-cli-viewer/src/templates"
)

type RuntimeConfigFlags struct {
	Test           *bool
	NoColors       *bool
	Template       *string
	ListTemplates  *bool
	InlineTemplate *string
	ShowFileChange *bool
	NoNewLine      *bool
	ParsedLineOnly *bool
}

var RuntimeConfig RuntimeConfigFlags

func init() {
	RuntimeConfig = RuntimeConfigFlags{
		Test:           flag.Bool("test", false, "Run against test data-set"),
		NoColors:       flag.Bool("c", false, "Disable colors"),
		Template:       flag.String("t", templates.DefaultTemplateName, "Template to use"),
		ListTemplates:  flag.Bool("T", false, "List available templates"),
		InlineTemplate: flag.String("i", "", "Inline template"),
		ShowFileChange: flag.Bool("f", false, "Show file change line"),
		NoNewLine:      flag.Bool("n", false, "Don't add empty lines between entries"),
		ParsedLineOnly: flag.Bool("p", false, "Parsed lines only"),
	}
	flag.Parse()
}
