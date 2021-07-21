package gampung

import (
	"context"
	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

type TemplateUrl struct {
	Url string
}

func GetTemplateUrl(q string, t string, p int) ([]TemplateUrl, int) {
	var templatesUrl []TemplateUrl
	var lastPage = p
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t},
		)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 100, Page: lastPage},
		Order: "desc",
	}

	templates, resp, err := client.Search.Code(ctx, q, opts)
	if  err != nil {
		return templatesUrl, 0
	}
	for _, url := range templates.CodeResults{
		if url.GetHTMLURL() == "" {
			continue
		}
		templatesUrl = append(templatesUrl, TemplateUrl{
			Url: convertURL(url.GetHTMLURL()),
		})
	}
	lastPage = resp.NextPage
	return templatesUrl, lastPage
}
