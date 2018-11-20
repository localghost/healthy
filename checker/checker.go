package checker

import (
	"github.com/localghost/healthy/utils"
	"time"
)

type Checker struct {
	checks map[string]Check
	metrics map[string]error
	request chan string
	responses map[string]chan error
}

type metric struct {
	name string
	value error
}

func New(checks interface{}) *Checker {
	result := &Checker{
		checks: make(map[string]Check),
		metrics: make(map[string]error),
		request: make(chan string),
		responses: make(map[string]chan error),
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
		c.responses[name] = make(chan error)
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
			case name := <- c.request:
				err, ok := c.metrics[name]
				if !ok {
					c.responses[name] <- utils.NewNoSuchCheckError(name)
				} else {
					c.responses[name] <- err
				}
			}
		}
	}()
}

func (c *Checker) Get(name string) error {
	c.request <- name
	return <-c.responses[name]
}

func (c* Checker) GetAll() error {
	for name := range c.checks {
		if err := c.Get(name); err != nil {
			return err
		}
	}
	return nil
}
