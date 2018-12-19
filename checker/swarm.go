package checker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/localghost/healthy/utils"
	"net/url"
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
	if _, err := url.Parse(c.Url); err != nil {
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

func (c *SwarmCheck) createClient() (dockerClient *client.Client, err error) {
	settings := []func(*client.Client) error{
		client.WithHost(c.Host),
		client.WithVersion(c.Version),
	}
	dockerClient, err = client.NewClientWithOpts(settings...)
}
