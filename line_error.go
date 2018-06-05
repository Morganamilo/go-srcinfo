package srcinfo

import (
	"fmt"
)

type LineError struct {
	LineNumber int
	Line       string
	ErrorStr   string
}

func (le LineError) Error() string {
	return fmt.Sprintf("Line %d: %s: %s", le.LineNumber, le.ErrorStr, le.Line)
}

func Error(LineNumber int, Line string, ErrorStr string) LineError {
	return LineError{
		LineNumber,
		Line,
		ErrorStr,
	}
}

func Errorf(LineNumber int, Line string, ErrorStr string, args ...interface{}) LineError {
	return LineError{
		LineNumber,
		Line,
		fmt.Sprintf(ErrorStr, args...),
	}
}
