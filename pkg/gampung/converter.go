package gampung

import "strings"

func convertURL(u string) string {
	rp := strings.NewReplacer("/blob/", "/", "github.com", "raw.githubusercontent.com")
	return rp.Replace(u)
}