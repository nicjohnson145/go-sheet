# go-sheet

Parses a sheet following the same schema as `example.yml` and displays the resulting character with
bonuses to attack/save/etc calculated.

### Building and Running

Requires go 1.16's `embed` feature

```
go build
./go-sheet /path/to/sheet.yml
```

If no sheet is given, the example sheet from this repo will be used

### Future Features

* Live reload when underlying character sheet changes
