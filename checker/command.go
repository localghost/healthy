package checker

import (
	"fmt"
	"github.com/localghost/healthy/utils"
	"log"
	"os/exec"
)

type CommandCheck struct {
	Command interface{}
	Shell string

	command []string
}

func NewCommandCheck() Check {
	return &CommandCheck{
		Shell: "sh",
	}
}

func (c *CommandCheck) Configure(options map[string]interface{}) error {
	if err := utils.Decode(options, c); err != nil {
		return err
	}

	switch c.Command.(type) {
	case string:
		c.command = []string{c.Shell, "-c", c.Command.(string)}
	case []interface{}:
		var command []string
		for _, arg := range c.Command.([]interface{}) {
			if argstr, ok := arg.(string); !ok {
				return fmt.Errorf("invalid command argument: %v", arg)
			} else {
				command = append(command, argstr)
			}
		}
		c.command = command
	default:
		return fmt.Errorf("command format is not supported")
	}
	return nil
}

func (c *CommandCheck) Run() error {
	log.Println("checking command:", c.command)
	if err := exec.Command(c.command[0], c.command[1:]...).Run(); err != nil {
		return err
	}
	return nil
}
