package main

import (
	"fmt"
	"time"

	"github.com/cloudflare/cloudflared/cmd/cloudflared/access"
	"github.com/urfave/cli/v2"
)

func SetupApp() *cli.App {
	return &cli.App{
		Name: 		"cloudflared",
		Usage: 		"Simplified cloudflared client to connect to tcp tunnels",
		Copyright:	fmt.Sprintf(
			`(c) %d Cloudflare Inc.
		Your installation of cloudflared software constitutes a symbol of your signature indicating that you accept
		the terms of the Apache License Version 2.0 (https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/license),
		Terms (https://www.cloudflare.com/terms/) and Privacy Policy (https://www.cloudflare.com/privacypolicy/).`,
			time.Now().Year(),
		),
		Flags:		access.Flags(),
		Commands:	access.Commands(),
	}
}

func RunTunnel(app *cli.App, args []string) {
	fmt.Printf("Starting tunnel\n")
	app.Run(args)
}
