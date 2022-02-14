package main

import (
	"github.com/Korazza/templay/cmd"
	"github.com/Korazza/templay/config"
	"github.com/Korazza/templay/logger"
)

func main() {
	var c config.Config

	if err := c.SetPwd(); err != nil {
		logger.Error.Fatal(err)
	}

	if err := c.Load(); err != nil {
		logger.Warning.Print(err)
	}

	if err := c.Validate(); err != nil {
		logger.Error.Fatal(err)
	}

	cmd.Execute(c)
}
