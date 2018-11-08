package checker

import (
	"fmt"
	"github.com/localghost/healthy/utils"
	"log"
	"net/http"
)

type Checker struct {
	checks map[string]check
	metrics map[string]error
	request chan request
}

type request struct {
	name string
	response chan error
}

type check struct {
	Type string
	Options map[string]interface{}
}

type metric struct {
	name string
	value error
}

func New(checks interface{}) *Checker {
	result := &Checker{
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
	c.checks = make(map[string]check)
	for name, check1 := range checks.(map[string]interface{}) {
		c.checks[name] = check{
			Type: (check1.(map[string]interface{}))["type"].(string),
			Options: check1.(map[string]interface{}),
		}
	}
	log.Println(c.checks)
}

func (c *Checker) startChecks() {
	receiver := make(chan metric)
	for name, ch := range c.checks {
		go func(name string, ch check) {
			if ch.Type == "http" {
				receiver <- metric{name, httpCheck(ch.Options["url"].(string))}
			}
		}(name, ch)
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
		name: name,
		response: response,
	}
	return <- response
}

func httpCheck(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("Invalid status code")
	}
	return nil
}
