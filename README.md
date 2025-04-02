# git-io

Generate and manage the static site at [git.huggins.io](https://git.huggins.io).

- The `cmd/sitegen` tool generates all the HTML for the redirect site and copies in any other assets. For repositories whose primary language is Go, it will also include the appropriate meta tag to tell pkg.go.dev how to handle the repository.
