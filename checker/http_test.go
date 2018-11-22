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

func TestMissingUrl(t *testing.T) {
	check := NewHttpCheck()
	if err := check.Configure(map[string]interface{}{}); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestInvalidUrlType(t *testing.T) {
	check := NewHttpCheck()
	if err := check.Configure(map[string]interface{}{"url": 42}); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestInvalidUrl(t *testing.T) {
	check := NewHttpCheck()
	if err := check.Configure(map[string]interface{}{"url": "invalid_url : 7777"}); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}
