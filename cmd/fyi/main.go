package main

import (
    "flag"
	"github.com/amirrezaask/fyi/config"
	"github.com/amirrezaask/fyi/pkg/http"
)

func main() {
    var cfgPath string
    flag.StringVar(&cfgPath)
	config.Init("")
	srv := http.NewEcho()
	err := srv.Start("")
	if err != nil {
		panic(err)
	}
}
