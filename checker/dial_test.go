package checker

import "testing"

func TestTcp(t *testing.T) {
	check := NewDialCheck()
	check.Configure(map[string]interface{}{
		"protocol": "tcp",
		"address": "8.8.8.8:53",
		"timeout": 5,
	})
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestUdp(t *testing.T) {
	check := NewDialCheck()
	check.Configure(map[string]interface{}{
		"protocol": "udp",
		"address": "8.8.4.4:53",
		"timeout": 5,
	})
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestDefaultProtocol(t *testing.T) {
	check := NewDialCheck()
	check.Configure(map[string]interface{}{
		"address": "8.8.8.8:53",
		"timeout": 5,
	})
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestDefaultTimeout(t *testing.T) {
	check := NewDialCheck()
	check.Configure(map[string]interface{}{
		"protocol": "udp",
		"address": "8.8.4.4:53",
	})
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidProtocol(t *testing.T) {
	check := NewDialCheck()
	check.Configure(map[string]interface{}{
		"protocol": "invalid",
		"address": "8.8.4.4:53",
		"timeout": 5,
	})
	if err := check.Run(); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

func TestInvalidAddress(t *testing.T) {
	check := NewDialCheck()
	check.Configure(map[string]interface{}{
		"protocol": "tcp",
		"address": "invalid-address and port",
		"timeout": 5,
	})
	if err := check.Run(); err == nil {
		t.Fatal("Expected check to fail but it succeeded.")
	}
}

//func TestMissingAddress(t *testing.T) {
//	check := NewDialCheck()
//	check.Configure(map[string]interface{}{
//		"protocol": "tcp",
//		"timeout": 5,
//	})
//	if err := check.Run(); err == nil {
//		t.Fatal("Expected check to fail but it succeeded.")
//	}
//}
