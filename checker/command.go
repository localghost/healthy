package checker

import (
	"fmt"
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
	var command []string
	switch options["command"].(type) {
	case string:
		command = []string{"sh", "-c", options["command"].(string)}
	case []string:
		command = options["command"].([]string)
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
