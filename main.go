
package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
)

var Version = "0.1.0"
var Name = "mysqlbeat"

func main() {
	nb := &MySQLbeat{}

	b := beat.NewBeat(Name, Version, nb)

	b.CommandLineSetup()

	b.LoadConfig()
	err := nb.Config(b)
	if err != nil {
		logp.Critical("Config error: %v", err)
		os.Exit(1)
	}

	b.Run()
}