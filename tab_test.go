package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestUntab(t *testing.T) {
	nstop = 2
	minspace = 2
	in := `
tab
	indented
		    space
		    aligned
`
	expected := `
tab
  indented
        space
        aligned
`
	out := convert(t, untab, in)
	if expected != out {
		t.Errorf("expected output:%sto be:%s", out, expected)
	}
}

func TestTab(t *testing.T) {
	nstop = 2
	minspace = 2
	in := `
tab
  indented
          space
          aligned
`
	expected := `
tab
	indented
					space
					aligned
`
	out := convert(t, tab, in)
	if expected != out {
		t.Errorf("expected output:%sto be:%s", out, expected)
	}
}

func convert(t *testing.T, cvt func(*bufio.Reader) error, in string) string {
	var outsb strings.Builder
	bout = bufio.NewWriter(&outsb)
	bin := bufio.NewReader(strings.NewReader(in))
	err := cvt(bin)
	if err != nil {
		t.Errorf("tab: %v", err)
	}
	bout.Flush()
	return outsb.String()
}
