package runner

import (
	"github.com/kankburhan/gampung/pkg/gampung"
	"github.com/projectdiscovery/gologger"
	"os"
	"path"
	"sync"
)

func New(options *Options){
	gologger.Print().Msgf("Searching Template From github...")
	var page = 1
	config := gampung.GetLastConfig()
	if config.LastPage != 0 {
		page = config.LastPage
	}

	var urls []gampung.TemplateUrl
	paging := make(chan int)
	var   wgf sync.WaitGroup
	for i := 0;  i < options.Concurrency; i++ {
		wgf.Add(1)
		go func() {
			for p := range paging {
				ls, lp := gampung.GetTemplateUrl(options.Query, options.Token, p)
				if lp > page {
					page = lp
				}
				if len(ls) > 0 {
					urls = append(urls, ls...)
				}
			}
			defer wgf.Done()
		}()
	}

	// loop page to  10 pages
	for i := page; i < page + 10; i++ {
		paging <- i
	}
	close(paging)
	wgf.Wait()

	gampung.SaveLastConfig(gampung.Configuration{LastPage: page})

	if len(urls) == 0 {
		gologger.Info().Msgf("No template found.")
		os.Exit(1)
	}

	gologger.Print().Msgf("Found %v Templates from github", len(urls))

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for url := range jobs{
				_, filename := path.Split(url)
				gampung.DownloadFile(options.Saveto, url, filename)
			}
			defer wg.Done()
		}()
	}

	for _, url := range urls{
		jobs <- url.Url
	}
	close(jobs)
	wg.Wait()
}