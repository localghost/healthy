package checker

import "testing"

func TestShellStyle(t *testing.T) {
	options := map[string]interface{}{
		"command": "echo foo",
	}
	check := NewCommandCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestExecStyle(t *testing.T) {
	options := map[string]interface{}{
		"command": []interface{}{"echo", "foo"},
	}
	check := NewCommandCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestShell(t *testing.T) {
	options := map[string]interface{}{
		"command": "[[ true ]] && echo",
		"shell": "bash",
	}
	check := NewCommandCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidCommand(t *testing.T) {
	options := map[string]interface{}{
		"command": 42,
	}
	check := NewCommandCheck()
	if err := check.Configure(options); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestMissingCommand(t *testing.T) {
	options := map[string]interface{}{}
	check := NewCommandCheck()
	if err := check.Configure(options); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}
