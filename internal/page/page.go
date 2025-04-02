// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package page

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

//go:embed style.css
var _css string

func stylesheet() gomponents.Node {
	var w strings.Builder
	m := minify.New()
	css.Minify(m, &w, strings.NewReader(_css), nil)

	return gomponents.Raw(w.String())
}

func RepositoryRedirect(repo github.Repository) gomponents.Node {
	head := []gomponents.Node{
		html.Meta(
			html.Name("go-import"),
			html.Content(fmt.Sprintf("git.huggins.io/%s git %s", repo.GetName(), repo.GetHTMLURL())),
		),
		gomponents.If(repo.GetLanguage() == "Go",
			html.Meta(
				html.Name("go-source"),
				// 4d63.com/vangen https://github.com/leighmcculloch/vangen https://github.com/leighmcculloch/vangen/tree/master{/dir} https://github.com/leighmcculloch/vangen/blob/master{/dir}/{file}#L{line}
				html.Content(fmt.Sprintf(
					"git.huggins.io/%s %s/tree/%s{/dir} %s/blob/%s{/dir}/{file}#L{line}",
					repo.GetName(), repo.GetHTMLURL(), repo.GetDefaultBranch(), repo.GetHTMLURL(), repo.GetDefaultBranch(),
				)),
			),
		),
	}

	return RedirectPage(head, repo.GetHTMLURL())
}

func UrlRedirect(url string) gomponents.Node {
	return RedirectPage(nil, url)
}

func RedirectPage(head []gomponents.Node, url string) gomponents.Node {
	return components.HTML5(components.HTML5Props{
		Language: "en-US",
		Title:    "Redirecting...",
		Head: append(
			head,
			gomponents.If(
				url != "#",
				html.Meta(gomponents.Attr("http-equiv", "refresh"), html.Content(fmt.Sprintf("0; url='%s'", url))),
			),
			html.StyleEl(stylesheet()),
		),
		Body: []gomponents.Node{
			html.Div(html.Class("container"),
				html.Span(html.A(html.Href(url), gomponents.Text("Click here if you are not redirected."))),
			),
		},
	})
}
