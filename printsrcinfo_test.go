package srcinfo

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func srcinfosEqual(t *testing.T, name, data, file string) {
	dataLines := strings.Split(data, "\n")
	fileLines := strings.Split(file, "\n")

data:
	for line, dataLine := range dataLines {
		trimmed := strings.TrimSpace(dataLine)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		for _, fileLine := range fileLines {
			if dataLine == fileLine {
				continue data
			}
		}

		t.Errorf("%s Line %d \"%s\" can not be found in source file", name, line, dataLine)
	}
}

func TestPrintSrcinfo(t *testing.T) {
	for _, name := range goodSrcinfos {
		path := filepath.Join(goodSrcinfoDir, name)
		srcinfo, err := ParseFile(path)
		if err != nil {
			t.Errorf("Error parsing %s: %s", name, err)
			continue
		}

		file, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("Unable to read file: %s: %s", path, err.Error())
			continue
		}

		srcinfosEqual(t, name, srcinfo.String(), string(file))
	}
}

func TestPrintSrcinfoEmpty(t *testing.T) {
	srcinfo := &Srcinfo{}

	str := srcinfo.String()
	if str != "" {
		t.Errorf("Empty srcinfo should generate empty string but gave: %s", str)
	}
}
