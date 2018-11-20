package checker

import (
	"fmt"
	"log"
	"net/http"
)

type HttpCheck struct {
	url string
}

func NewHttpCheck() Check {
	return &HttpCheck{}
}

func (c *HttpCheck) Configure(options map[string]interface{}) error {
	var ok bool
	if c.url, ok = options["url"].(string); !ok {
		return fmt.Errorf("invalid or missing url")
	}
	return nil
}

func (c *HttpCheck) Run() error {
	log.Println("checking url:", c.url)
	response, err := http.Get(c.url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("invalid status code")
	}
	return nil
}
