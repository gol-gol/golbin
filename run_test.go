package golbin

import (
	"strings"
	"testing"
)

func TestConsoleRun(t *testing.T) {
	kon := Console{Command: "echo 'Hello Shell!'"}
	kon.Run()
	if strings.Trim(kon.StdOutput, " ") != "Hello Shell!\n" {
		t.Errorf("FAILED to echo required string. Got: '%s'", kon.StdOutput)
	}
}

func TestConsoleRunFailure(t *testing.T) {
	kon := Console{Command: "no-cmd-echo 'Hello Shell!'"}
	kon.Run()

	hasError := false
	for _, out := range strings.Fields(kon.StdOutput) {
		if out == "Error:" {
			hasError = true
		}
	}
	if !hasError {
		t.Errorf("FAILED to echo required string. Got: '%s'", kon.StdOutput)
	}
}

func TestExecOutput(t *testing.T) {
	out := ExecOutput("echo 'Hello Shell!'")
	if strings.Trim(out, " ") != "Hello Shell!\n" {
		t.Errorf("FAILED to echo required string. Got: '%s'", out)
	}
}

func TestExec(t *testing.T) {
	out, err := Exec("echo 'Hello Shell!'")
	if strings.Trim(out, " ") != "Hello Shell!\n" || err != nil {
		t.Errorf("FAILED to echo required string. Got: '%s' & '%s'", out, err.Error())
	}

	out, err = Exec("should-not-echo 'Hello Shell!'")
	if strings.Trim(out, " ") != "" || err == nil {
		t.Errorf("FAILED to echo required string. Got: '%s' & '%s'", out, err.Error())
	}
}

func TestExecWithEnv(t *testing.T) {
	out, err := ExecWithEnv("bash -c 'echo $TEMP_GOLBIN_VAR'", map[string]string{"TEMP_GOLBIN_VAR": "Hey!"})
	if strings.Trim(out, " ") != "Hey!\n" || err != nil {
		t.Errorf("FAILED to echo required string. Got: '%s' & '%s'", out, err.Error())
	}
}
