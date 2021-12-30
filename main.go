package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/cmd"
)

func main() {
	ctx := log.Logger.WithContext(context.Background())
	cmd.Execute(ctx)
}
