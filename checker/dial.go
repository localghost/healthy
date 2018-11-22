package checker

import (
	"fmt"
	"github.com/localghost/healthy/utils"
	"log"
	"net"
	"time"
)

type DialCheck struct {
	Protocol string
	Address string
	Timeout time.Duration
}

var protocols = map[string]struct{} {
	"tcp": {},
	"udp": {},
	"ip": {},
	"unix": {},
}

func NewDialCheck() Check {
	return &DialCheck{
		Protocol: "tcp",
		Timeout: 10 * time.Second,
	}
}

func (d* DialCheck) Configure(options map[string]interface{}) error {
	if err := utils.Decode(options, d); err != nil {
		return err
	}
	if _, ok := protocols[d.Protocol]; !ok {
		return fmt.Errorf("unsupported protocol %s", d.Protocol)
	}
	if d.Address == "" {
		return fmt.Errorf("address is not set")
	}
	return nil
}

func (d* DialCheck) Run() error {
	conn, err := net.DialTimeout(d.Protocol, d.Address, d.Timeout)
	if err != nil {
		return err
	}
	if err = conn.Close(); err != nil {
		log.Println(err)
	}
	return nil
}
