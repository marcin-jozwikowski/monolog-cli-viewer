package templates

import "text/template"

const DefaultTemplateName = "normal"

var templates = map[string]string{
	DefaultTemplateName: `{{._color}}{{._level}}:{{.channel}}{{._colorReset}}{{"\t"}}{{._datetime}}{{"\t"}}{{._color}}{{.message}}{{._colorReset}}{{"\r\n"}}{{._context}}{{"\r\n"}}`,
	"full":              `{{._color}}{{._level}}:{{.channel}}{{._colorReset}}{{"\t"}}{{._datetime}}{{"\t"}}{{._color}}{{.message}}{{._colorReset}}{{"\r\n"}}{{._contextPretty}}{{"\r\n"}}`,
	"min":               `{{._color}}{{._level}}:{{.channel}}{{"\t"}}{{._datetime}}{{"\t"}}{{.message}}{{._colorReset}}`,
}

func GetTemplatateByName(name string) (*template.Template, error) {
	templateString, exists := templates[name]
	if exists == false {
		templateString = templates[DefaultTemplateName]
	}

	return GetTemplatateFromString(templateString)
}

func GetTemplatateFromString(templateString string) (*template.Template, error) {
	return template.New("item").Parse(templateString)
}

func GetTemplateNames() []string {
	keys := make([]string, 0, len(templates))
	for k := range templates {
		keys = append(keys, k)
	}

	return keys
}
