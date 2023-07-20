package apiserver

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("APISERVER_CONFIG_PATH", "../../../configs/apiserver.toml")
}

func TestAPIServer_HandleHello(t *testing.T) {
	apiServerConfig, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	apiServer := New(apiServerConfig)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	apiServer.handleHello().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "Hello")
}
