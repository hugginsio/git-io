package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
		log.Println("Rendering ", repo.GetName(), "(language:", repo.GetLanguage(), ")")
		html := page.GenericRedirect(repo.GetHTMLURL(), strings.ToLower(repo.GetLanguage()) == "go")

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

	log.Println("Repository pages generated")
	log.Println("Copying individual assets")

	robots := "User-agent: *\nDisallow: /"
	f, err := os.Create("_output/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	if _, err := f.WriteString(robots); err != nil {
		log.Fatalln(err)
	}
}
