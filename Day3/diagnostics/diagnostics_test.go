package main

import (
	_ "embed"
	"testing"
)

//go:embed diagnostics_report_test.txt
var testInput string

func TestDiagnostics(t *testing.T) {
	t.Run("Running Diagnostics Report", func(t *testing.T) {
		TestDiagnosticsReport := DiagnosticsReport{}
		TestDiagnosticsReport.GenerateParameters(testInput)
		got := TestDiagnosticsReport.Product()
		want := 198

		if got != want {
			t.Errorf("got: %d, want: %d", got, want)
		}
	})
}

func TestLifeSupport(t *testing.T) {
	t.Run("Running Life Support Report", func(t *testing.T) {
		TestDiagnosticsReport := DiagnosticsReport{}
		TestDiagnosticsReport.GenerateLifeSupportParameters(testInput)
		got := TestDiagnosticsReport.LifeSupportProduct()
		want := 230

		if got != want {
			t.Errorf("got: %q, expected: %q", got, want)
		}
	})
}
