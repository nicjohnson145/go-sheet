# /usr/bin/env bash

set -euo pipefail

fyne-cross linux -name go-sheet-linux-amd64 -arch amd64
fyne-cross windows -arch amd64

TAG=$(svu next)

gh release create $TAG
gh release upload $TAG fyne-cross/bin/linux-amd64/go-sheet-linux-amd64 fyne-cross/bin/windows-amd64/go-sheet.exe
