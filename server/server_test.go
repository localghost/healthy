package server

import (
	"github.com/golang/mock/gomock"
	"github.com/localghost/healthy/checker"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
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

func TestSingleCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	checker := checker.NewMockChecker(ctrl)
	checker.EXPECT().Get(gomock.Eq("foo")).Return(nil)

	router := New(checker).newRouter()

	req, err := http.NewRequest("GET", "/v1/check/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("Expected request to succeed but it failed with: %s", err)
	}
}
