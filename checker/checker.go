package checker

import (
	"fmt"
	"github.com/localghost/healthy/utils"
	"log"
	"net/http"
	"time"
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

var checkFunctions = map[string]func (map[string]interface{}) error {
	"http": httpCheck,
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
}

func (c *Checker) startChecks() {
	receiver := make(chan metric)
	for name, ch := range c.checks {
		go func(name string, ch check, checkFunction func (map[string]interface{}) error) {
			for {
				select {
				case <- time.After(10 * time.Second):
					receiver <- metric{name, checkFunction(ch.Options)}
				}
			}
		}(name, ch, checkFunctions[ch.Type])
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

func httpCheck(options map[string]interface{}) error {
	log.Println("checking url:", options["url"].(string))
	response, err := http.Get(options["url"].(string))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("Invalid status code")
	}
	return nil
}
