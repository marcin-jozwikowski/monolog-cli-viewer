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

const result_line_1__some_test = "NOTICE:default\t2023-11-09 14:29:06\tSome test message\r\n{\"session\":{\"id\":\"bq2fk4i3nhkgbj4eua964g5r63\"},\"user\":{\"id\":1}}"
const result_line_2__checking_support = "DEBUG:security\t2023-11-14 00:37:26\tChecking support on authenticator.\r\n{\"authenticator\":\"App\\\\Security\\\\AppAuthenticator\",\"firewall_name\":\"main\"}"
const result_line_3__unparsed = "2023-10-23 11:03:16: [9a4e77e9afa8] [ERROR] [Whatever] Login Error"
const result_line_4__file_change = "==> some/path/to\\file.log <=="
const result_line_5__user_logged = "INFO:default\t2023-10-23 11:07:47\tUser logged in\r\n{\"user\":{\"id\":\"54767261-98c6-4a57-9064-0d35fd06d1fc\"}}"
const result_line_6__route_matched = "INFO:request\t2023-12-29 10:26:40\tMatched route \"api_login\".\r\n{\"method\":\"POST\",\"request_uri\":\"http://localhost/api/v1/login\",\"route\":\"api_login\",\"route_parameters\":{\"_controller\":\"App\\\\User\\\\Infrastructure\\\\Controller\\\\ApiLoginController::index\",\"_route\":\"api_login\"}}"
const result_line__empty = ""

func TestDefault(t *testing.T) {
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           false,
		ShowFileChangeLine:  false,
		ShowParsedLinesOnly: false,
		Template:            templ,
	})

	runTests(t, "TestDefault", []string{
		result_line_1__some_test + "\r\n\n",
		result_line_2__checking_support + "\r\n\n",
		result_line_3__unparsed + "\n\n",
		result_line__empty,
		result_line_5__user_logged + "\r\n\n",
		result_line_6__route_matched + "\r\n\n",
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
		result_line_1__some_test + "\r\n",
		result_line_2__checking_support + "\r\n",
		result_line_3__unparsed + "\n",
		result_line__empty,
		result_line_5__user_logged + "\r\n",
		result_line_6__route_matched + "\r\n",
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
		result_line_1__some_test + "\r\n\n",
		result_line_2__checking_support + "\r\n\n",
		result_line_3__unparsed + "\n\n",
		result_line_4__file_change + "\n\n",
		result_line_5__user_logged + "\r\n\n",
		result_line_6__route_matched + "\r\n\n",
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
		result_line_1__some_test + "\r\n\n",
		result_line_2__checking_support + "\r\n\n",
		result_line__empty,
		result_line__empty,
		result_line_5__user_logged + "\r\n\n",
		result_line_6__route_matched + "\r\n\n",
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
		result_line_1__some_test + "\r\n",
		result_line_2__checking_support + "\r\n",
		result_line__empty,
		result_line_4__file_change + "\n",
		result_line_5__user_logged + "\r\n",
		result_line_6__route_matched + "\r\n",
	})
}

func runTests(t *testing.T, name string, expected []string) {
	count := len(expected)
	for id, line := range data.GetTestData() {
		logLineItem := viewer.InitLogLine(line)
		result := logLineItem.GetFormattedString()

		if id == count {
			t.Errorf("Response %d out of bounds on %s. Did not expect: '%q'", id, name, result)
		}

		if result != expected[id] {
			t.Errorf("Invalid string value on test '%s' item '%d'. Expected\r\n'%q', got\r\n'%q'", name, id, expected[id], result)
		}
	}
}
