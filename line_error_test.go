package srcinfo

import (
	"testing"
)

const lineNumber int = 5
const line string = "pkgbase ="
const errorStr string = "Value is empty"

func TestLineError(t *testing.T) {
	err := Error(lineNumber, line, errorStr)

	if err.LineNumber != lineNumber {
		t.Errorf("Line number should be %d but was %d", lineNumber, err.LineNumber)
	}

	if err.Line != line {
		t.Errorf("Line should be \"%s\" but was \"%s\"", line, err.Line)
	}

	if err.ErrorStr != errorStr {
		t.Errorf("ErrorStr should be \"%s\" but was \"%s\"", errorStr, err.ErrorStr)
	}

	t.Logf("error: %#v generated message: %s", err, err.Error())
}

func TestLineErrorf(t *testing.T) {
	err := Errorf(lineNumber, line, "%s", errorStr)

	if err.LineNumber != lineNumber {
		t.Errorf("Line number should be %d but was %d", lineNumber, err.LineNumber)
	}

	if err.Line != line {
		t.Errorf("Line should be \"%s\" but was \"%s\"", line, err.Line)
	}

	if err.ErrorStr != errorStr {
		t.Errorf("ErrorStr should be \"%s\" but was \"%s\"", errorStr, err.ErrorStr)
	}

	t.Logf("error: %#v generated message: %s", err, err.Error())
}
