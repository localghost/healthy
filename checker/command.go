package checker

import (
	"fmt"
	"github.com/localghost/healthy/utils"
	"log"
	"os/exec"
)

type CommandCheck struct {
	command []string
}

func NewCommandCheck() Check {
	return &CommandCheck{}
}

func (c *CommandCheck) Configure(options map[string]interface{}) error {
	var ok bool

	var shell string
	if shell, ok = utils.WithDefault(options, "shell", "sh").(string); !ok {
		return fmt.Errorf("invalid shell provided")
	}

	var command []string
	switch options["command"].(type) {
	case string:
		command = []string{shell, "-c", options["command"].(string)}
	case []interface{}:
		for _, arg := range options["command"].([]interface{}) {
			if argstr, ok := arg.(string); !ok {
				return fmt.Errorf("invalid command argument: %v", arg)
			} else {
				command = append(command, argstr)
			}
		}
	default:
		return fmt.Errorf("command format is not supported")
	}
	c.command = command
	return nil
}

func (c *CommandCheck) Run() error {
	log.Println("checking command:", c.command)
	if err := exec.Command(c.command[0], c.command[1:]...).Run(); err != nil {
		return err
	}
	return nil
}
