package checker

import (
	"github.com/localghost/healthy/utils"
	"time"
)

type Checker struct {
	checks map[string]Check
	metrics map[string]error
	request chan request
}

type request struct {
	name string
	response chan error
}

type metric struct {
	name string
	value error
}

func New(checks interface{}) *Checker {
	result := &Checker{
		checks: make(map[string]Check),
		metrics: make(map[string]error),
		request: make(chan request),
	}
	result.parseChecks(checks)
	return result
}

func (c *Checker) Start() {
	c.startChecks()
}

func (c *Checker) parseChecks(checks interface{}) {
	for name, check := range checks.(map[string]interface{}) {
		ctype := (check.(map[string]interface{}))["type"].(string)
		options := check.(map[string]interface{})
		c.checks[name] = registry.CreateAndConfigure(ctype, options)
	}
}

func (c *Checker) startChecks() {
	receiver := make(chan metric)
	for name, check := range c.checks {
		go func(name string, check Check) {
			for {
				select {
				case <- time.After(10 * time.Second):
					receiver <- metric{name, check.Run()}
				}
			}
		}(name, check)
	}
	go func() {
		for {
			select {
			case m := <-receiver:
				c.metrics[m.name] = m.value
			case r := <- c.request:
				err, ok := c.metrics[r.name]
				if !ok {
					r.response <- utils.NewNoSuchCheckError(r.name)
				} else {
					r.response <- err
				}
			}
		}
	}()
}

func (c *Checker) Check(name string) error {
	response := make(chan error)
	c.request <- request{
		name:     name,
		response: response,
	}
	return <-response
}
