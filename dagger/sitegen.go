// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/git-io/internal/dagger"
)

// Run cmd/sitegen
func (m *GitIo) Sitegen(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Directory {
	return m.buildEnv(ctx, source).
		WithExec([]string{"go", "run", "./cmd/sitegen/"}).
		Directory("/go/src/_output/")
}
