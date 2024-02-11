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
const result_line_7__doctrine_exception = "ERROR:request\t2023-12-31 11:18:53\tUncaught PHP Exception Symfony\\Component\\HttpKernel\\Exception\\HttpException: \"This value should be of type unknown. This value should not be blank. This value should not be blank.\" at RequestPayloadValueResolver.php line 127\r\n{\"exception\":\"[object] (Symfony\\\\Component\\\\HttpKernel\\\\Exception\\\\HttpException(code: 0): This value should be of type unknown.\\nThis value should not be blank.\\nThis value should not be blank. at /var/www/vendor/symfony/http-kernel/Controller/ArgumentResolver/RequestPayloadValueResolver.php:127)\\n[previous exception] [object] (Symfony\\\\Component\\\\Validator\\\\Exception\\\\ValidationFailedException(code: 0): :\\n    This value should be of type unknown.\\nObject(App\\\\Scooter\\\\Infrastructure\\\\Request\\\\UpdateLocation\\\\UpdateLocationRequest).latitude:\\n    This value should not be blank. (code c1051bb4-d103-4f74-8988-acbcafc7fdc3)\\nObject(App\\\\Scooter\\\\Infrastructure\\\\Request\\\\UpdateLocation\\\\UpdateLocationRequest).longitude:\\n    This value should not be blank. (code c1051bb4-d103-4f74-8988-acbcafc7fdc3)\\n at /var/www/vendor/symfony/http-kernel/Controller/ArgumentResolver/RequestPayloadValueResolver.php:127)\"}"
const result_line_8__doctrine_query = "DEBUG:doctrine\t2023-12-29 11:35:33\tExecuting statement: SELECT t0.id AS id_1, t0.email AS email_2, t0.roles AS roles_3, t0.password AS password_4 FROM user_entity t0 WHERE t0.id = ? (parameters: array{\"1\":1}, types: array{\"1\":1})\r\n{\"params\":{\"1\":1},\"sql\":\"SELECT t0.id AS id_1, t0.email AS email_2, t0.roles AS roles_3, t0.password AS password_4 FROM user_entity t0 WHERE t0.id = ?\",\"types\":{\"1\":1}}"
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
		result_line_7__doctrine_exception + "\r\n\n",
		result_line_8__doctrine_query + "\r\n\n",
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
		result_line_7__doctrine_exception + "\r\n",
		result_line_8__doctrine_query + "\r\n",
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
		result_line_7__doctrine_exception + "\r\n\n",
		result_line_8__doctrine_query + "\r\n\n",
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
		result_line_7__doctrine_exception + "\r\n\n",
		result_line_8__doctrine_query + "\r\n\n",
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
		result_line_7__doctrine_exception + "\r\n",
		result_line_8__doctrine_query + "\r\n",
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
