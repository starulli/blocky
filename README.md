# Blocky
Converts pixel art image files into SVGs for the web.

## Usage
Using `blocky` is pretty simple: you give it a PNG file and it writes the SVG
data to standard output.
```
$ blocky [-debug] [-keepInvisible] FILE
```
The flags `blocky` accepts mean the following:

| Flag                   | Default | Explanation                                  |
| ---------------------- | ------- | -------------------------------------------- |
| `-debug`               | Off     | Enables debug mode output                    |
| `-keepInvisible`       | Off     | Write elements that have `0x00` alpha values |
| `-exclude=#RRGGBB[AA]` | Off     | Excludes the given colour from output. Alpha is `0xFF` when unspecified. |

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

The test case is a `455B`, `25x20` PNG file; these numbers will most likely differ
based on PNG file, but it should still be a decent ballpark.

| Version | Description              | Size (B) |  ΔSize | GZ (B) |  ΔGZ | ~Ratio (%) |
| ------- | ------------------------ | -------: | -----: | -----: | ---: | ---------: |
| v0.1    | Pixel-for-pixel output   |   33,412 |      – |  1,646 |    – | 20:1 (95%) |
| v0.2    | Alpha & colour exclusion |   12,487 |   ↓63% |    831 | ↓49% | 15:1 (93%) |

Currently, the generated SVG is `1.8x` larger than the source image.
