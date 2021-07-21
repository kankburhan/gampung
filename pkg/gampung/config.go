package gampung

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
)

const (
	gampungpath = "/.gampung/"
	gampungconfig = "config.json"
)

type Configuration struct {
	LastPage int
}

func GetLastConfig() Configuration{
	var config Configuration
	data, err := ioutil.ReadFile(userHomeDir() + gampungpath + gampungconfig)
	if err == nil {
		err = json.Unmarshal(data, &config)
		if err == nil {
			return config
		}
	}
	return config
}

func SaveLastConfig(c Configuration)  {
	os.MkdirAll(
		userHomeDir()+gampungpath,
		0755,
	)
	config, _ := json.Marshal(c)
	_ = ioutil.WriteFile(
			userHomeDir() + gampungpath + gampungconfig,
			config,
			0644,
	)
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}