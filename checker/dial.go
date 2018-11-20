package checker

import (
	"fmt"
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

var protocols = map[string]struct{} {
	"tcp": {},
	"udp": {},
	"ip": {},
	"unix": {},
}

func NewDialCheck() Check {
	return &DialCheck{}
}

func (d* DialCheck) Configure(options map[string]interface{}) error {
	var ok bool
	var address, protocol string
	var timeout int

	if address, ok = options["address"].(string); !ok {
		return fmt.Errorf("address not defined")
	}
	if protocol, ok = utils.WithDefault(options, "protocol", "tcp").(string); !ok {
		return fmt.Errorf("invalid protocol format")
	}
	if _, ok = protocols[protocol]; !ok {
		return fmt.Errorf("unsupported protocol %s", protocol)
	}
	if timeout, ok = utils.WithDefault(options, "timeout", 10).(int); !ok {
		return fmt.Errorf("invalid timeout %v", timeout)
	}

	d.address = address
	d.protocol = protocol
	d.timeout = time.Duration(timeout) * time.Second

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
