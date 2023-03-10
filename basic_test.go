package golbin

import "testing"

func TestIsSystemCmd(t *testing.T) {
	if !IsSystemCmd("go") {
		t.Error("FAILED to identify system cmd: 'go'")
	}
	if IsSystemCmd("should-not-exist-go") {
		t.Error("FAILED; identified system cmd: 'should-not-exist-go'")
	}
}
