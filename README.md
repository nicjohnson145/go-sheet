# go-sheet

Parses a sheet following the same schema as `example.yml` and displays the resulting character with
bonuses to attack/save/etc calculated.

### Building and Running

There are pre-compiled binaries for Linux, Windows, & Mac
[here](https://github.com/nicjohnson145/go-sheet/releases/latest). Or optionally clone this repo and
run `go build`

### Dependencies

If using the `go build` approach, ensure the necessary OS packages are installed as detailed
[here](https://github.com/go-gl/glfw)

To run go-sheet, simple supply the path to the character yaml on the command line. Go-sheet will
watch this file for changes and update whenever a change is detected

```
./go-sheet /path/to/sheet.yml
```

If no sheet is given, the example sheet from this repo will be used.

