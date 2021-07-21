package gampung

import (
	"github.com/projectdiscovery/gologger"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		gologger.Error().Msgf("Failed Download %s", filename)
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath+filename)
	if err != nil {
		gologger.Error().Msgf("Failed Write %s", filename)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err == nil {
		gologger.Info().Msgf("Success Download %s", filename)
	}
	return err
}