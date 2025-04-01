package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	if err := os.Mkdir("_output", 0755); err != nil {
		log.Fatalln(err)
	}

	for _, repo := range repos {
		log.Println("Processing repository:", repo.GetName(), "(language:", repo.GetLanguage(), ")")
		// TODO: language specfic config for golang meta
		html := page.GenericRedirect(repo.GetHTMLURL())

		if err := os.Mkdir(fmt.Sprintf("_output/%s", repo.GetName()), 0755); err != nil {
			log.Fatalln(err)
		}

		f, err := os.Create(fmt.Sprintf("_output/%s/index.html", repo.GetName()))
		if err != nil {
			log.Fatalln(err)
		}

		defer f.Close()

		if err := html.Render(f); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("HTML generated")
}
