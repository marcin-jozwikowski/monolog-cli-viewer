# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.11.0] - 2024-01-24

### Added

- `-p` flag to display only lines that were correctly parsed. No other log formats shall pass.
- Tests in the github workflows

## [0.10.1] - 2024-01-23

### Added

- Check for CHANGELOG in the workflow before new release

## [0.10.0] - 2024-01-23

### Added

- `-f` flag to display the file change lines coming from `tail`
- `-n` flag to remove empty lines between log entries

## [0.9.0] - 2023-12-05

### Added

- Automated builds for OSX and Windows
- Automated release with github pipeline

## [0.1.0] - 2023-12-02

### Added

- Reading the inputs from pipeline
- Parsing JSON to variables
- Formatting output with text/template
- 3 build-in templates
- Level-based colors
- Disabling colors with -c option
- Selecting template with -t option
- Inline template with -i option
