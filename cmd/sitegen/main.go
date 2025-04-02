// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"fmt"
	"log"

	"git.huggins.io/git-io/internal/fsio"
	"git.huggins.io/git-io/internal/page"
	"github.com/google/go-github/github"
)

func main() {
	log.Println("Fetching repositories from forge")
	gh := github.NewClient(nil)
	repos, _, err := gh.Repositories.List(context.Background(), "hugginsio", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fsio.Delete("_output")
	fsio.Directory("_output", 0755)

	for _, repo := range repos {
		log.Println("Rendering ", repo.GetName(), "(language:", repo.GetLanguage(), ")")

		html := page.RepositoryRedirect(*repo)
		fsio.Directory(fmt.Sprintf("_output/%s", repo.GetName()), 0755)

		// typical
		f := fsio.File(fmt.Sprintf("_output/%s.html", repo.GetName()))
		defer f.Close()

		if err := html.Render(f); err != nil {
			log.Fatalln(err)
		}

		// pretty
		f = fsio.File(fmt.Sprintf("_output/%s/index.html", repo.GetName()))
		defer f.Close()

		if err := html.Render(f); err != nil {
			log.Fatalln(err)
		}

	}

	log.Println("Repository pages generated")
	log.Println("Copying individual assets")

	robots := "User-agent: *\nDisallow: /"
	f := fsio.File("_output/robots.txt")
	defer f.Close()

	if _, err := f.WriteString(robots); err != nil {
		log.Fatalln(err)
	}

	index := page.UrlRedirect("https://github.com/hugginsio")
	f = fsio.File("_output/index.html")
	defer f.Close()

	if err := index.Render(f); err != nil {
		log.Fatalln(err)
	}
}
