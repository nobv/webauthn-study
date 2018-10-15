package server

import (
	"net/http/httptest"
	"testing"

	"github.com/Nobv/webauthn/server"
)

func TestNewServer(t *testing.T) {

	addr := "localhost:8080"

	server := server.NewServer(addr)

	req := httptest.NewRequest("GET", "/".nil)

}
