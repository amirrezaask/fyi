package config

import (
    "github.com/golobby/config/v2"
    "github.com/golobby/config/v2/feeder"

)

var C *config.Config

func Init(path string) error {
	var err error
    C, err = config.New(feeder.Json{Path: path})
	if err != nil {
		return err
	}
	return nil
}
