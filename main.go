//go:generate pkger
package main

import (
	"context"

	"github.com/markbates/pkger"
	"github.com/rs/zerolog/log"
	"gitlab.com/stackvista/stackstate-cli2/cmd"
)

func main() {
	// Instruct pkger to include the entire /static directory
	_ = pkger.Include("/static")
	ctx := log.Logger.WithContext(context.Background())
	cmd.Execute(ctx)
}
