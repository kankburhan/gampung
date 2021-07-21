package main

import runner "github.com/kankburhan/gampung/internal"

func main(){
	option := runner.ParseOptions()
	runner.New(option)
}