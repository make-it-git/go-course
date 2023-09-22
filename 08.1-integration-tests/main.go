package main

import (
	"context"

	"08-tests/logic"
)

func main() {
	client := logic.GetClient("localhost:6379")
	ctx := context.Background()
	logic.SetValue(ctx, client, "some value")
}
