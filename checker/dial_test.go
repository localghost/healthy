package checker

import "testing"

func TestTcp(t *testing.T) {
	options := map[string]interface{}{
		"protocol": "tcp",
		"address": "8.8.8.8:53",
		"interval": "5s",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestUdp(t *testing.T) {
	options := map[string]interface{}{
		"protocol": "udp",
		"address": "8.8.4.4:53",
		"interval": "5s",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestDefaultProtocol(t *testing.T) {
	options := map[string]interface{}{
		"address": "8.8.8.8:53",
		"interval": "5s",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestDefaultTimeout(t *testing.T) {
	options := map[string]interface{}{
		"protocol": "udp",
		"address": "8.8.4.4:53",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidProtocol(t *testing.T) {
	options := map[string]interface{}{
		"protocol": "invalid",
		"address": "8.8.4.4:53",
		"interval": "5s",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestInvalidAddress(t *testing.T) {
	options := map[string]interface{}{
		"protocol": "tcp",
		"address": "invalid-address and port",
		"interval": "5s",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestMissingAddress(t *testing.T) {
	options := map[string]interface{}{
		"protocol": "tcp",
		"interval": "5s",
	}
	check := NewDialCheck()
	if err := check.Configure(options); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestMissingOptions(t *testing.T) {
	check := NewHttpCheck()
	if err := check.Configure(map[string]interface{}{}); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}
