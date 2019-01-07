package checker

import (
	"context"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"testing"
)

func TestRunningSwarm(t *testing.T) {
	if !isSwarmEnabled() {
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

func TestCheckFailureOnMissingService(t *testing.T) {
	if !isSwarmEnabled() {
		t.Skip("Swarm not enabled")
	}
	options := map[string]interface{}{
		"services": []string{"missing.service"},
	}
	check := NewSwarmCheck()
	if err := check.Configure(options); err != nil {
		t.Fatal(err)
	}
	if err := check.Run(); err == nil {
		t.Fatal("Expected test to fail but it succeeded")
	}
}

func isSwarmEnabled() bool {
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
