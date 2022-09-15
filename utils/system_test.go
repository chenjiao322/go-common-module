package utils

import "testing"

func Test_getGID(t *testing.T) {
	if GetGID() != GetGID() {
		t.Error("different gid")
	}
}
