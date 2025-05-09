// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"dagger/git-io/internal/dagger"
)

type GitIo struct{}

// Return the build environment container.
func (m *GitIo) buildEnv(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	return dag.Container().
		From("golang:1.24-alpine").
		WithDirectory("/go/src/", source).
		WithMountedCache("/go/pkg/mod/", dag.CacheVolume("go-mod-124")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithWorkdir("/go/src/")
}
