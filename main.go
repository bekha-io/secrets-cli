package main

import (
	"log"
	"os"

	"github.com/bekha-io/secrets/commands"
	"github.com/bekha-io/secrets/db"
	"github.com/urfave/cli"
)

var app = cli.NewApp()


func applicationInfo() {
	app.Name = "Secrets CLI"
	app.Usage = "A simple command line interface for storing and accessing your secrets"
	app.Author = "Bekhruz Iskandarzoda (github.com/bekha-io)"
	app.Version = "1.0.0"
}


func main() {
	applicationInfo()
	commands.AddCommands_Groups(app)
	commands.AddCommands_Secrets(app)

	db.SetupDatabase()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Cannot run application: %v", err)
	}
}