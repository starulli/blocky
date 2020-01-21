# Blocky
Converts pixel art image files into SVGs for the web.

## Usage
Using `blocky` is pretty simple:
```
$ blocky [-debug] FILE
```
All output is written to `stdout`, so you may redirect it to save output.

The flags `blocky` accepts mean the following:

| Flag     | Default | Explanation               |
| -------- | ------- | ------------------------- |
| `-debug` | Off     | Enables debug mode output |

### Example
To run in debug mode, while converting `artwork.png` into `logo.svg`:
```
$ blocky -debug artwork.png > logo.svg
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

### Debug Mode
When the `-debug` flag is passed, the resulting SVG will have stroke styling
embedded into it so that the rectangle outputs can be distinguished:
```
$ make build
$ ./blocky -debug image.png > image.svg
```

## SVG Resouce Sizes
This tool was created to turn pixel art PNG files into SVGs for website use, so
it's important to keep track of how the generated SVG elements will affect page
sizes.

The test case is a `165B`, `9x9` PNG file; these numbers will most likely differ
based on PNG file, but it should still be a decent ballpark.

| Version | Description             | Size (B) |  ΔSize | GZ (B) |  ΔGZ | ~Ratio (%) |
| ------- | ----------------------- | -------: | -----: | -----: | ---: | ---------: |
| v0.1    | pixel-for-pixel output  |    4,806 |      – |    420 |    – | 11:1 (91%) |

Currently, the generated SVG is `2.5x` larger than the source image.
