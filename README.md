# Blocky
Converts pixel art image files into SVGs for the web.

## Usage
Using `blocky` is pretty simple: you give it a PNG file and it writes the SVG
data to standard output.
```
$ blocky [-debug] [-keepInvisible] [-exclude=#RRGGBB[AA]] [-optimize=pixels|rects] FILE
```
The flags `blocky` accepts mean the following:

| Flag                   | Default | Explanation                                  |
| ---------------------- | ------- | -------------------------------------------- |
| `-debug`               | Off     | Enables debug mode output                    |
| `-keepInvisible`       | Off     | Write elements that have `0x00` alpha values |
| `-exclude=#RRGGBB[AA]` | Off     | Excludes the given colour from output. Alpha is `0xFF` when unspecified. |
| `-optimize`            | Off     | Optimize rectangles in the SVG               |

### Optimize Flag Values
The optimize flag supports the following modes:

| Mode     | Description                                 |
| -------- | ------------------------------------------- |
| Off      | No optimization. Defaults to `pixels`       |
| `pixels` | Emits individual pixels                     |
| `rects`  | Emits rectangles for contiguous colours     |

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
| v0.1    | Pixel-for-pixel output   |   33,411 |      – |  1,649 |    – | 20:1 (95%) |
| v0.2    | Alpha & colour exclusion |   12,486 |   ↓63% |    835 | ↓50% | 15:1 (93%) |
| v0.3    | Rectangle optimization   |    6,610 |   ↓47% |    649 | ↓22% | 10:1 (90%) |

Currently, the generated SVG is `1.4x` larger than the source image. There are
other opportunities for further size optimizations:

- Organizing output to take better advantage of compression
- Looking for the largest NxM rectangle instead of Nx1 and 1xN
- Finding colour islands and pushing them into a higher z-index

However, the size increase over the source image is acceptable to me for now,
since the source image isn't actually delivered to the clients. If we look at
linear scaling of the source image to some realistic resolutions, the current
SVG size is adequate:

| WxH     |  Size (B) |
| ------- |  -------: |
| SVG     |       649 |
| 25x20   |       455 |
| 75x60   |     1,365 |
| 150x120 |     2,730 |
