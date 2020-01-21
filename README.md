# Blocky
Converts pixel art image files into SVGs for the web.

## Usage
Basic usage is simple; all output is written to `stdout`, so just redirect
it to save the SVG:
```
$ blocky image.png > image.svg
```

## Development
Requires Go 1.13 to build and has no external dependencies.

### Testing
The default action for testing is to run all tests:
```
$ make test
```
Set the `SPEC` variable to target a specific package during testing:
```
$ make test SPEC=./svg
```
