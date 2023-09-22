package tests

import (
	"context"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"

	"08-tests/logic"
)

func TestSetValue(t *testing.T) {
	ctx := context.Background()

	redisContainer, err := redis.RunContainer(ctx,
		testcontainers.WithImage("docker.io/redis:7"),
	)
	if err != nil {
		panic(err)
	}

	dsn, err := redisContainer.ConnectionString(ctx)
	if err != nil {
		panic(err)
	}
	dsn = strings.Replace(dsn, "redis://", "", 1)
	client := logic.GetClient(dsn)
	result := logic.SetValue(context.Background(), client, "my test value")
	if result != "my test value" {
		t.Fail()
	}

	// Clean up the container
	defer func() {
		if err := redisContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
}
