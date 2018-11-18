package checker

import (
	"github.com/localghost/healthy/utils"
	"log"
	"net"
	"time"
)

type DialCheck struct {
	protocol string
	address string
	timeout time.Duration
}

func NewDialCheck() Check {
	return &DialCheck{}
}

func (d* DialCheck) Configure(options map[string]interface{}) error {
	d.address = options["address"].(string)
	d.protocol = utils.WithDefault(options, "protocol", "tcp").(string)
	d.timeout = time.Duration(utils.WithDefault(options, "timeout", 10).(int)) * time.Second
	return nil
}

func (d* DialCheck) Run() error {
	conn, err := net.DialTimeout(d.protocol, d.address, d.timeout)
	if err != nil {
		return err
	}
	if err = conn.Close(); err != nil {
		log.Println(err)
	}
	return nil
}
