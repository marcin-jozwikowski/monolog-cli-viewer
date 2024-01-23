package main

import (
	"bufio"
	"fmt"
	"monolog-cli-viewer/src/cli"
	"monolog-cli-viewer/src/colors"
	"monolog-cli-viewer/src/templates"
	"monolog-cli-viewer/src/viewer"
	"os"
	"strings"
	"text/template"
)

func main() {
	fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.

	if *cli.RuntimeConfig.ListTemplates == true {
		fmt.Println("Possible templates are: " + strings.Join(templates.GetTemplateNames(), ", "))
		return
	}

	colors.SetEnabled(!*cli.RuntimeConfig.NoColors) // enable or disable the colors based on CLI flag
	viewer.SetSettings(viewer.Settings{
		NoNewLine:           *cli.RuntimeConfig.NoNewLine,      // don't add empty lines
		ShowFileChangeLine:  *cli.RuntimeConfig.ShowFileChange, // show file change from tail
		ShowParsedLinesOnly: *cli.RuntimeConfig.ParsedLineOnly, // don't show unparsed lines
		Template:            initiateTemplate(),                // get template from the CLI params, or the default one
	})
	if *cli.RuntimeConfig.Test == true {
		// test values @todo - put those to tests
		values := []string{
			`{"message":"Some test message","context":{"user":{"id":1},"session":{"id":"bq2fk4i3nhkgbj4eua964g5r63"}},"level":"NOTICE","channel":"default","timestamp":"1699540146122"}`,
			`{"message":"Checking support on authenticator.","context":{"firewall_name":"main","authenticator":"App\\Security\\AppAuthenticator"},"level":100,"level_name":"DEBUG","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}`,
			`2023-10-23 11:03:16: [9a4e77e9afa8] [ERROR] [Whatever] Login Error`,
			`==> some/path/to\file.log <==`,
			`[2023-10-23T11:07:47.038324+00:00] default.INFO: User logged in {"user":{"id":"54767261-98c6-4a57-9064-0d35fd06d1fc"}} []`,
		}

		for _, line := range values {
			logLineItem := viewer.InitLogLine(line)
			fmt.Print(logLineItem.GetFormattedString())
		}
		return
	}

	if (fi.Mode() & os.ModeCharDevice) == 0 { // we're in the pipe
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() { // wait for a line of text
			line := scanner.Text()

			logLineItem := viewer.InitLogLine(line)     // init the logItem from line string
			fmt.Print(logLineItem.GetFormattedString()) // format it and print it
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("reading files is still not implemented. Use `cat file.log | mcv`")
	}
}

func initiateTemplate() *template.Template {
	var t *template.Template
	var e error
	if *cli.RuntimeConfig.InlineTemplate != "" {
		t, e = templates.GetTemplatateFromString(*cli.RuntimeConfig.InlineTemplate) // get template based on content provided in CLI flag
	} else {
		t, e = templates.GetTemplatateByName(*cli.RuntimeConfig.Template) // get template based on name provided in CLI flag
	}
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	return t
}
