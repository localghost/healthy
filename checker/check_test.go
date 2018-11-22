package checker

import (
	"errors"
	"log"
	"testing"
)

type TestCheckMock struct {
	configured bool
	data map[string]interface{}

	failConfigure bool
	failRun bool
}
func NewTestCheckMock() Check {
	return &TestCheckMock{}
}
func FailingTestCheckMockFactory(failConfigure, failRun bool) func () Check {
	return func () Check {
		return &TestCheckMock{
			failConfigure: failConfigure,
			failRun: failRun,
		}
	}
}
func (m *TestCheckMock) Configure(options map[string]interface{}) error {
	if m.failConfigure {
		return errors.New("")
	}
	m.configured = true
	m.data = options
	return nil
}
func (m *TestCheckMock) Run() error {
	if m.failRun {
		return errors.New("")
	}
	return nil
}

func TestCheckRegistry_CreateCheck(t *testing.T) {
	registry := NewCheckRegistry()
	registry.Add("mock", NewTestCheckMock)

	check, err := registry.CreateAndConfigure("mock", map[string]interface{}{"foo": "bar"})
	if err != nil {
		log.Fatalf("Configuration failed %s", err)
	}
	if !check.(*TestCheckMock).configured {
		log.Fatal("Get not configured.")
	}
	if foo, ok := check.(*TestCheckMock).data["foo"]; !ok || foo != "bar" {
		log.Fatalf("Configured key is not valid, expected 'bar' got '%s'", foo)
	}
}

func TestCheckRegistry_CreateNonExisting(t *testing.T) {
	registry := NewCheckRegistry()

	_, err := registry.CreateAndConfigure("mock", map[string]interface{}{})
	if err == nil {
		log.Fatalf("Expected creating non-existing check to fail but it succeeded")
	}
}

func TestCheckRegistry_FailConfigure(t *testing.T) {
	registry := NewCheckRegistry()
	registry.Add("mock", FailingTestCheckMockFactory(true, false))

	_, err := registry.CreateAndConfigure("mock", map[string]interface{}{})
	if err == nil {
		log.Fatalf("Expected configuring check to fail but it succeeded")
	}
}

func TestCheckRegistry_FailRun(t *testing.T) {
	registry := NewCheckRegistry()
	registry.Add("mock", FailingTestCheckMockFactory(false, true))

	check, err := registry.CreateAndConfigure("mock", map[string]interface{}{})
	if err != nil {
		log.Fatalf("Configuration failed %s", err)
	}
	if err = check.Run(); err == nil {
		log.Fatal("Expected run check to fail but it succeeded")
	}
}
