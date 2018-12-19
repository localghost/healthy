package checker

import (
	"github.com/docker/docker/client"
	"github.com/localghost/healthy/utils"
)

func NewSwarmCheck() Check {
	return &SwarmCheck{
		Host: "unix:///var/run/docker.sock",
		Version: "auto",
	}
}

type SwarmCheck struct {
	Host string
	Version string

	client *client.Client
}

func (c *SwarmCheck) Configure(options map[string]interface{}) error {
	if err := utils.Decode(options, c); err != nil {
		return err
	}
	return nil
}

func (c *SwarmCheck) Run() (err error) {
	if c.client == nil {
		if c.client, err = c.createClient(); err != nil {
			return
		}
	}
	// ...
	return
}

func (c *SwarmCheck) createClient() (*client.Client, error) {
	settings := []func(*client.Client) error{
		client.WithHost(c.Host),
		client.WithVersion(c.Version),
	}
	return client.NewClientWithOpts(settings...)
}
