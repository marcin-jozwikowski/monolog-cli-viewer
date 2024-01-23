# Monolog CLI viewer

### A simple tool to format monolog output into more readable form

## Installation 

1. Download latest binary for your OS from [releases](https://github.com/marcin-jozwikowski/monolog-cli-viewer/releases)
1. Place the executable in your system PATH and ensure it has runing priviliges

## Usage 

Add `mcv` or `mcv.exe` into the pipeline that outputs the monolog logs. Some examples:

* `cat monolog.log | grep 'user deleted' | mcv` - for searching a specific entry
* `tail -f monolog.log | mcv` - for live preview of a file
* `docker logs php-nginx -f | mcv.exe` - for a live preview of docker logs

### Parameters

The following parameters are supported 

| Param | Description | Default | 
| --- | --- | --- |
| -t | Name of the template to use | normal |
| -i | Template body to be used instead of predefined ones. See [Templates](#templates) for more. |

As well as the following flags
| Flag | Description |
| --- | --- |
| -T | List all templates available to use |
| -c | Disable colors in the output |
| -f | Show `==> file.name <==` lines from `tail` command | 
| -n | Don't add empty lines between entries | 
| -test | Run the config against a predefined data-set |

### Templates

The application uses GoLang [textTemplates](https://pkg.go.dev/text/template) with all fields from incoming monolog JSON available at root level, and traversable by dot-notation, and starting with a `.` (dot). Example: `{{.message}}`, and `{{.channel}}`. See [Additional Fields](#additional-fields) for more.

Whenever a field in the tamplete will contain some sub-values (i.e. `{{._context}}`) it will display a map (`map[id:123]`). The full path of `{{._context.id}}` will just display the value `123`.

When the requested path will not be present at all, it will display `<no value>`

### Additional Fields

Additional fields are added to the values from monolog to be available in the template. They are:

* `_colors` - colors and text formatting options. This value contains the following sub-values:
  *  `red`, `green`, `yellow`, `blue`, `purple`, `cyan`, `white` for colors
  * `bold`, `underline`, `strike`, `italic` for formatting
  * `reset` - for resetting the text output back to normal
* `_color` - one of the available colors based on monolog log level
  * DEBUG - white
  * INFO - cyan
  * NOTICE - green
  * WARNING - yellow
  * ERROR - purple
  * CRITICAL - red
  * ALERT - red
  * EMERGENCY - red
* `_colorR` - alias for `_colors.reset`
* `_level` - string representation of the int values in monolog `level`
* `_dateTime` - a `Y-m-d H:i:s` formatted value passed in `timestamp` or `datetime` fields of the monolog JSON
* `_context` and `_extra` - string representations of the JSON object passed in `context` and `extra` monolog fields
* `_contextPretty` and `_extraPretty` pretty-formatted versions of the fields

### Behind the scenes

MCV parses each line of the incoming data and tries to display it using a provided template.

As of now the [JSON](https://github.com/Seldaek/monolog/blob/main/src/Monolog/Formatter/JsonFormatter.php) and [Line](https://github.com/Seldaek/monolog/blob/main/src/Monolog/Formatter/LineFormatter.php) formatters are supported with their default config and default [message structure](https://github.com/Seldaek/monolog/blob/main/doc/message-structure.md).

As the program is designed to be used in pipeline it also recognizes the file name line from `tail`. It is ignored by default.