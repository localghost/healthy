package checker

import (
	"github.com/localghost/healthy/utils"
)

type Checker interface {
	Get(name string) error
	GetAll() error
}

type CheckerImpl struct {
	tasks     map[string]*Task
	metrics   map[string]error
	request   chan string
	responses map[string]chan error
}

type Spec struct {
	Type string
}

type Task struct {
	spec  Spec
	check Check
}

func NewChecker(checks interface{}) (Checker, error) {
	result := &CheckerImpl{
		tasks:     make(map[string]*Task),
		metrics:   make(map[string]error),
		request:   make(chan string),
		responses: make(map[string]chan error),
	}
	if err := result.parseChecks(checks); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *CheckerImpl) parseChecks(checks interface{}) error {
	var specs = make(map[string]Spec)
	if err := utils.Decode(checks, &specs); err != nil {
		return err
	}

	for name, check := range checks.(map[string]interface{}) {
		options := check.(map[string]interface{})
		if check, err := registry.CreateAndConfigure(specs[name].Type, options); err != nil {
			return err
		} else {
			c.tasks[name] = &Task{
				spec:  specs[name],
				check: check,
			}
			c.responses[name] = make(chan error)
		}
	}
	return nil
}

func (c *CheckerImpl) Get(name string) error {
	if _, ok := c.tasks[name]; !ok {
		return utils.NewNoSuchCheckError(name)
	}
	return c.tasks[name].check.Run()
}

func (c*CheckerImpl) GetAll() error {
	for name := range c.tasks {
		if err := c.Get(name); err != nil {
			return err
		}
	}
	return nil
}
