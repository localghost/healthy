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
	Services []string
	MinReplicas uint64

	client *client.Client
}

func (c *SwarmCheck) Configure(options map[string]interface{}) (err error) {
	if err = utils.Decode(options, c); err != nil {
		return
	}
	c.client, err = client.NewClientWithOpts(client.WithHost(c.Host)) // , client.WithVersion(c.Version))
	return
}

func (c *SwarmCheck) Run() (err error) {
	var services []swarm.Service
	if services, err = c.getServices(); err != nil {
		return
	}
	var tasksByService map[string][]swarm.Task
	if tasksByService, err = c.getTasksByService(); err != nil {
		return
	}
	for _, service := range services {
		if service.Spec.Mode.Replicated != nil {
			if err = c.checkReplicatedService(service, tasksByService[service.ID]); err != nil {
				return
			}
		} else {
			if err = c.checkGlobalService(service, tasksByService[service.ID]); err != nil {
				return
			}
		}
	}
	return
}

func (c *SwarmCheck) getServices() (services []swarm.Service, err error) {
	if len(c.Services) == 0 {
		return c.client.ServiceList(context.Background(), types.ServiceListOptions{})
	}
	var kv []filters.KeyValuePair
	for _, service := range c.Services {
		kv = append(kv, filters.Arg("name", service))
	}
	return c.client.ServiceList(context.Background(), types.ServiceListOptions{Filters: filters.NewArgs(kv...)})
}

func (c *SwarmCheck) getTasksByService() (map[string][]swarm.Task, error) {
	if tasks, err := c.client.TaskList(context.Background(), types.TaskListOptions{}); err != nil {
		return nil, err
	} else {
		tasksByService := make(map[string][]swarm.Task)
		for _, task := range tasks {
			tasksByService[task.ServiceID] = append(tasksByService[task.ServiceID], task)
		}
		return tasksByService, nil
	}
}

func (c *SwarmCheck) countRunningTasks(tasks []swarm.Task) (count uint64) {
	for _, task := range tasks {
		if task.Status.State == swarm.TaskStateRunning {
			count++
		}
	}
	return
}

func (c *SwarmCheck) checkReplicatedService(service swarm.Service, tasks []swarm.Task) error {
	minReplicas := *service.Spec.Mode.Replicated.Replicas
	if c.MinReplicas > 0 {
		minReplicas = c.MinReplicas
	}
	if c.countRunningTasks(tasks) < minReplicas {
		return fmt.Errorf("too few running tasks for service %s", service.Spec.Name)
	}
	return nil
}

func (c *SwarmCheck) checkGlobalService(service swarm.Service, tasks []swarm.Task) error {
	minReplicas := uint64(1) // get number of active nodes
	if c.MinReplicas > 0 {
		minReplicas = c.MinReplicas
	}
	if c.countRunningTasks(tasks) < minReplicas {
		return fmt.Errorf("no running tasks for global service %s", service.Spec.Name)
	}
	return nil
}
