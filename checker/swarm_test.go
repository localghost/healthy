package checker

import (
	"context"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"testing"
)

func TestRunningSwarm(t *testing.T) {
	if !hasSwarmEnabled() {
		t.Skip("Swarm not enabled")
	}
	options := map[string]interface{}{}
	check := NewSwarmCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err != nil {
		t.Fatal(err)
	}
}

func hasSwarmEnabled() bool {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return false
	}
	info, err := cli.Info(context.Background())
	if err != nil {
		return false
	}
	return info.Swarm.LocalNodeState == swarm.LocalNodeStateActive
}
