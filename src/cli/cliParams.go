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
}

var RuntimeConfig RuntimeConfigFlags

func init() {
	RuntimeConfig = RuntimeConfigFlags{
		Test:           flag.Bool("test", false, "Run against test data-set"),
		NoColors:       flag.Bool("c", false, "Disable colors"),
		Template:       flag.String("t", templates.DefaultTemplateName, "Template to use"),
		ListTemplates:  flag.Bool("T", false, "List available templates"),
		InlineTemplate: flag.String("i", "", "Inline template"),
	}
	flag.Parse()
}
