package checker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
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

func (c *SwarmCheck) Configure(options map[string]interface{}) (err error) {
	if err = utils.Decode(options, c); err == nil {
		return
	}
	c.client, err = c.createClient()
	return
}

func (c *SwarmCheck) Run() (err error) {
	var tasks []swarm.Task
	var filter = filters.NewArgs(filters.Arg("desired-state", "running"))
	if tasks, err = c.client.TaskList(context.Background(), types.TaskListOptions{Filters: filter}); err != nil {
		return
	}
	for _, task := range tasks {
		if task.Status.State != task.DesiredState {
			return fmt.Errorf("expected task %s in state %s but got %s", task.ID, task.DesiredState, task.Status.State)
		}
	}
	return
}

func (c *SwarmCheck) createClient() (*client.Client, error) {
	settings := []func(*client.Client) error{
		client.WithHost(c.Host),
		client.WithVersion(c.Version),
	}
	return client.NewClientWithOpts(settings...)
}
