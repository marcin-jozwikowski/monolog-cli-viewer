package main

import (
	"monolog-cli-viewer/src/data"
	"monolog-cli-viewer/src/templates"
	"monolog-cli-viewer/src/viewer"
	"testing"
	"text/template"
)

var templ *template.Template

func init() {
	templ, _ = templates.GetTemplatateByName(templates.DefaultTemplateName)
}

func TestDefault(t *testing.T) {
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           false,
		ShowFileChangeLine:  false,
		ShowParsedLinesOnly: false,
		Template:            templ,
	})

	runTests(t, "TestDefault", []string{
		"NOTICE:default\t2023-11-09 14:29:06\tSome test message\r\n{\"session\":{\"id\":\"bq2fk4i3nhkgbj4eua964g5r63\"},\"user\":{\"id\":1}}\r\n\n",
		"DEBUG:security\t2023-11-14 00:37:26\tChecking support on authenticator.\r\n{\"authenticator\":\"App\\\\Security\\\\AppAuthenticator\",\"firewall_name\":\"main\"}\r\n\n",
		"2023-10-23 11:03:16: [9a4e77e9afa8] [ERROR] [Whatever] Login Error\n\n",
		"",
		"INFO:default\t2023-10-23 11:07:47\tUser logged in \r\n{\"user\":{\"id\":\"54767261-98c6-4a57-9064-0d35fd06d1fc\"}}\r\n\n",
	})
}

func TestNoNewLine(t *testing.T) {
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           true,
		ShowFileChangeLine:  false,
		ShowParsedLinesOnly: false,
		Template:            templ,
	})

	runTests(t, "TestNoNewLine", []string{
		"NOTICE:default\t2023-11-09 14:29:06\tSome test message\r\n{\"session\":{\"id\":\"bq2fk4i3nhkgbj4eua964g5r63\"},\"user\":{\"id\":1}}\r\n",
		"DEBUG:security\t2023-11-14 00:37:26\tChecking support on authenticator.\r\n{\"authenticator\":\"App\\\\Security\\\\AppAuthenticator\",\"firewall_name\":\"main\"}\r\n",
		"2023-10-23 11:03:16: [9a4e77e9afa8] [ERROR] [Whatever] Login Error\n",
		"",
		"INFO:default\t2023-10-23 11:07:47\tUser logged in \r\n{\"user\":{\"id\":\"54767261-98c6-4a57-9064-0d35fd06d1fc\"}}\r\n",
	})
}

func TestShowFileChange(t *testing.T) {
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           false,
		ShowFileChangeLine:  true,
		ShowParsedLinesOnly: false,
		Template:            templ,
	})

	runTests(t, "TestShowFileChange", []string{
		"NOTICE:default\t2023-11-09 14:29:06\tSome test message\r\n{\"session\":{\"id\":\"bq2fk4i3nhkgbj4eua964g5r63\"},\"user\":{\"id\":1}}\r\n\n",
		"DEBUG:security\t2023-11-14 00:37:26\tChecking support on authenticator.\r\n{\"authenticator\":\"App\\\\Security\\\\AppAuthenticator\",\"firewall_name\":\"main\"}\r\n\n",
		"2023-10-23 11:03:16: [9a4e77e9afa8] [ERROR] [Whatever] Login Error\n\n",
		"==> some/path/to\\file.log <==\n\n",
		"INFO:default\t2023-10-23 11:07:47\tUser logged in \r\n{\"user\":{\"id\":\"54767261-98c6-4a57-9064-0d35fd06d1fc\"}}\r\n\n",
	})
}

func TestParsedLinesOnly(t *testing.T) {
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           false,
		ShowFileChangeLine:  false,
		ShowParsedLinesOnly: true,
		Template:            templ,
	})

	runTests(t, "TestParsedLinesOnly", []string{
		"NOTICE:default\t2023-11-09 14:29:06\tSome test message\r\n{\"session\":{\"id\":\"bq2fk4i3nhkgbj4eua964g5r63\"},\"user\":{\"id\":1}}\r\n\n",
		"DEBUG:security\t2023-11-14 00:37:26\tChecking support on authenticator.\r\n{\"authenticator\":\"App\\\\Security\\\\AppAuthenticator\",\"firewall_name\":\"main\"}\r\n\n",
		"",
		"",
		"INFO:default\t2023-10-23 11:07:47\tUser logged in \r\n{\"user\":{\"id\":\"54767261-98c6-4a57-9064-0d35fd06d1fc\"}}\r\n\n",
	})
}

func TestNoNewLineFileChangeParsedOnly(t *testing.T) {
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           true,
		ShowFileChangeLine:  true,
		ShowParsedLinesOnly: true,
		Template:            templ,
	})

	runTests(t, "TestNoNewLineFileChangeParsedOnly", []string{
		"NOTICE:default\t2023-11-09 14:29:06\tSome test message\r\n{\"session\":{\"id\":\"bq2fk4i3nhkgbj4eua964g5r63\"},\"user\":{\"id\":1}}\r\n",
		"DEBUG:security\t2023-11-14 00:37:26\tChecking support on authenticator.\r\n{\"authenticator\":\"App\\\\Security\\\\AppAuthenticator\",\"firewall_name\":\"main\"}\r\n",
		"",
		"==> some/path/to\\file.log <==\n",
		"INFO:default\t2023-10-23 11:07:47\tUser logged in \r\n{\"user\":{\"id\":\"54767261-98c6-4a57-9064-0d35fd06d1fc\"}}\r\n",
	})
}

func runTests(t *testing.T, name string, expected []string) {
	for id, line := range data.GetTestData() {
		logLineItem := viewer.InitLogLine(line)
		result := logLineItem.GetFormattedString()

		if result != expected[id] {
			t.Errorf("Invalid string value on '%s' item '%d'. Expected '%q', got '%q'", "testDefault", id, expected[id], result)
		}
	}
}
