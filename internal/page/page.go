// Copyright 2025 Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package page

import (
	_ "embed"
	"fmt"
	"strings"

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

func GenericRedirect(url string, repository string, goImport bool) gomponents.Node {
	return components.HTML5(components.HTML5Props{
		Language: "en-US",
		Title:    "Redirecting...",
		Head: []gomponents.Node{
			html.StyleEl(stylesheet()),
			// NOTE: only use meta refresh for known URLs.
			gomponents.If(
				url != "#",
				html.Meta(gomponents.Attr("http-equiv", "refresh"), html.Content(fmt.Sprintf("0; url='%s'", url))),
			),
			gomponents.If(
				goImport,
				html.Meta(gomponents.Attr("go-import"), html.Content(fmt.Sprintf("git.huggins.io/%s git %s", repository, url))),
			),
		},
		Body: []gomponents.Node{
			html.Div(html.Class("container"),
				html.Span(html.A(html.Href(url), gomponents.Text("Click here if you are not redirected."))),
			),
		},
	})
}
