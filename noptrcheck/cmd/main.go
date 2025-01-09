package main

import (
	"noptrcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(noptrcheck.Analyzer)
}
