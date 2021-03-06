package main

import (
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/integration/checker"
	"github.com/go-check/check"
)

// Test case for #22244
func (s *DockerSuite) TestAuthApi(c *check.C) {
	testRequires(c, Network)
	config := types.AuthConfig{
		Username: "no-user",
		Password: "no-password",
	}

	expected := "Get https://registry-1.docker.io/v2/: unauthorized: incorrect username or password"
	status, body, err := sockRequest("POST", "/auth", config)
	c.Assert(err, check.IsNil)
	c.Assert(status, check.Equals, http.StatusUnauthorized)
	msg := getErrorMessage(c, body)
	c.Assert(msg, checker.Contains, expected, check.Commentf("Expected: %v, got: %v", expected, msg))
}
