package parser

import (
	"strings"
	"testing"
	"time"
)

func TestParseSeconds(t *testing.T) {
	tests := []struct {
		in string
		d  time.Duration
	}{
		{"", 0},
		{"4", 4 * time.Second},
		{"0.1", 100 * time.Millisecond},
		{"0.050", 50 * time.Millisecond},
		{"2.003", 2*time.Second + 3*time.Millisecond},
	}

	for _, test := range tests {
		d := parseSeconds(test.in)
		if d != test.d {
			t.Errorf("parseSeconds(%q) == %v, want %v\n", test.in, d, test.d)
		}
	}
}

func TestParseNanoseconds(t *testing.T) {
	tests := []struct {
		in string
		d  time.Duration
	}{
		{"", 0},
		{"0.1", 0 * time.Nanosecond},
		{"0.9", 0 * time.Nanosecond},
		{"4", 4 * time.Nanosecond},
		{"5000", 5 * time.Microsecond},
		{"2000003", 2*time.Millisecond + 3*time.Nanosecond},
	}

	for _, test := range tests {
		d := parseNanoseconds(test.in)
		if d != test.d {
			t.Errorf("parseSeconds(%q) == %v, want %v\n", test.in, d, test.d)
		}
	}
}

func TestParseMessage(t *testing.T) {
	const test = `
=== RUN   TestExpandVisibleOriginalTargets
    TestExpandVisibleOriginalTargets: state_test.go:56:
        	Error Trace:	state_test.go:56
        	Error:      	Not equal:
        	            	expected: core.BuildLabels{core.BuildLabel{PackageName:"src/core", Name:"_target1#zip", Subrepo:""}, core.BuildLabel{PackageName:"src/core", Name:"target1", Subrepo:""}}
        	            	actual  : core.BuildLabels{core.BuildLabel{PackageName:"src/core", Name:"target1", Subrepo:""}}

        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -1,3 +1,2 @@
        	            	-(core.BuildLabels) (len=2) {
        	            	- (core.BuildLabel) //src/core:_target1#zip,
        	            	+(core.BuildLabels) (len=1) {
        	            	  (core.BuildLabel) //src/core:target1
        	Test:       	TestExpandVisibleOriginalTargets
--- FAIL: TestExpandVisibleOriginalTargets (0.00s)
FAIL
`
	report, err := Parse(strings.NewReader(test), "test")
	if err != nil {
		t.Errorf("Unexpected error parsing test: %s", err)
	}
	if len(report.Packages[0].Tests[0].Output) != 16 {
		t.Errorf("Unexpected output from test: was %d lines, should be 16", len(report.Packages[0].Tests[0].Output))
	}
}
