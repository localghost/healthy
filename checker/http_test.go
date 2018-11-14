package checker

import "testing"

func TestExistingUrl(t *testing.T) {
	check := NewHttpCheck()
	check.Configure(map[string]interface{}{
		"url": "http://google.com",
	})
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestNonExistingUrl(t *testing.T) {
	check := NewHttpCheck()
	check.Configure(map[string]interface{}{
		"url": "http://some-non-existing-domain-probably.xxx",
	})
	if err := check.Run(); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}
