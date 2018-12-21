package checker

import "fmt"

type Check interface {
	Configure(options map[string]interface{}) error

	Run() error
}

type CheckRegistry struct {
	providers map[string]func () Check
}

func NewCheckRegistry() *CheckRegistry {
	return &CheckRegistry{
		providers: make(map[string]func () Check),
	}
}

func (r *CheckRegistry) Add(name string, provider func() Check) {
	r.providers[name] = provider
}

func (r *CheckRegistry) CreateAndConfigure(name string, options map[string]interface{}) (Check, error) {
	if provider, ok := r.providers[name]; !ok {
		return nil, fmt.Errorf("check type %s is not supported", name)
	} else {
		check := provider()
		if err := check.Configure(options); err != nil {
			return nil, err
		}
		return check, nil
	}
}

var registry = NewCheckRegistry()

func init() {
	registry.Add("http", NewHttpCheck)
	registry.Add("dial", NewDialCheck)
	registry.Add("command" , NewCommandCheck)
	registry.Add("swarm", NewSwarmCheck)
}
