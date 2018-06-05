package srcinfo

import (
	"testing"
)

func TestSrcinfo(t *testing.T) {
	srcinfo, err := ParseSrcinfo("srcinfo")
	if err != nil {
		t.Error(err)
	} else {
		t.Error("\n" + srcinfo.String())
	}
}
