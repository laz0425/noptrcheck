package noptrcheck_test

import (
	"noptrcheck"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, noptrcheck.Analyzer, "noptrcheck")
}
