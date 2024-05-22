package testutils

import (
	"context"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func GetTestContainer() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Env: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "postgres",
		},
		HostConfigModifier: func(hc *container.HostConfig) { hc.NetworkMode = "host" },
		ExposedPorts:       []string{"5432/tcp"},
		Image:              "postgres:16.3",
		WaitingFor: wait.ForExec([]string{"pg_isready"}).
			WithPollInterval(2 * time.Second).
			WithExitCodeMatcher(func(exitCode int) bool {
				return exitCode == 0
			}),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
