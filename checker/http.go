package checker

import (
	"fmt"
	"github.com/localghost/healthy/utils"
	"log"
	"net/http"
	"net/url"
)

type HttpCheck struct {
	Url string
}

func NewHttpCheck() Check {
	return &HttpCheck{}
}

func (c *HttpCheck) Configure(options map[string]interface{}) error {
	if err := utils.Decode(options, c); err != nil {
		return err
	}
	if c.Url == "" {
		return fmt.Errorf("URL not set")
	}
	if _, err := url.Parse(c.Url); err != nil {
		return err
	}
	return nil
}

func (c *HttpCheck) Run() error {
	log.Println("checking url:", c.Url)
	response, err := http.Get(c.Url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("invalid status code")
	}
	return nil
}
