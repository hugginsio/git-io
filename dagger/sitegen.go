package main

import (
	"context"
	"dagger/git-io/internal/dagger"
)

func (m *GitIo) Sitegen(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Directory {
	return m.BuildEnv(ctx, source).
		WithExec([]string{"go", "run", "./cmd/sitegen/"}).
		Directory("/go/src/_output/")
}
