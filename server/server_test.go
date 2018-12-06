package server

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/localghost/healthy/checker"
	"github.com/localghost/healthy/utils"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNoCheckerProvided(t *testing.T) {
	defer func() {
		recover()
	}()

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
		t.Fatalf("Expected successful response but got response with code %d", rr.Code)
	}
}

func TestNoSuchCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	checker := checker.NewMockChecker(ctrl)
	checker.EXPECT().Get(gomock.Eq("foo")).Return(utils.NewNoSuchCheckError("foo"))

	router := New(checker).newRouter()

	req, err := http.NewRequest("GET", "/v1/check/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Fatalf("Expected response with code %d but got %d", http.StatusNotFound, rr.Code)
	}
}

func TestCheckFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	checker := checker.NewMockChecker(ctrl)
	checker.EXPECT().Get(gomock.Eq("foo")).Return(fmt.Errorf("check failed"))

	router := New(checker).newRouter()

	req, err := http.NewRequest("GET", "/v1/check/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusExpectationFailed {
		t.Fatalf("Expected response with code %d but got %d", http.StatusExpectationFailed, rr.Code)
	}
	if rr.Body.String() != "check failed" {
		t.Fatalf("Expected response body 'check failed' but got %s", rr.Body.String())
	}
}

func TestStatusCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	checker := checker.NewMockChecker(ctrl)
	checker.EXPECT().GetAll().Return(nil)

	router := New(checker).newRouter()

	req, err := http.NewRequest("GET", "/v1/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("Expected successful response but got response with code %d", rr.Code)
	}
}

func TestStatusCheckFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	checker := checker.NewMockChecker(ctrl)
	checker.EXPECT().GetAll().Return(fmt.Errorf("status failed"))

	router := New(checker).newRouter()

	req, err := http.NewRequest("GET", "/v1/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusExpectationFailed {
		t.Fatalf("Expected response with code %d but got %d", http.StatusExpectationFailed, rr.Code)
	}
	if rr.Body.String() != "status failed" {
		t.Fatalf("Expected response body 'status failed' but got %s", rr.Body.String())
	}
}

func TestMain(m *testing.M) {
	viper.Set("server.listen_on", "localhost:1234")
	os.Exit(m.Run())
}
