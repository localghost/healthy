package server

import (
	"github.com/spf13/viper"
	"testing"
)

func TestNoCheckerProvided(t *testing.T) {
	defer func() {
		recover()
	}()

	viper.Set("server.listen_on", "localhost:1234")
	New(nil)
	t.Fatal("server should panic on nil checker")
}
