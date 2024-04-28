package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "rnd"
	app.Usage = "Random String & Number Generator"
	app.Version = "0.0.0" //TODO: Update to use CI to update version

	// Flags
	lengthFlag := cli.IntFlag{
		Name:    "length",
		Aliases: []string{"l"},
		Usage:   "length",
	}
	prefixFlag := cli.StringFlag{
		Name:    "prefix",
		Aliases: []string{"p"},
		Usage:   "prefix",
	}
	bitFlag := cli.IntFlag{
		Name:    "bits",
		Aliases: []string{"b"},
		Usage:   "bits",
	}
	byteFlag := cli.IntFlag{
		Name:    "bytes",
		Aliases: []string{"B"},
		Usage:   "bytes",
	}

	// Commands
	app.Commands = []*cli.Command{
		{
			Name:    "base62",
			Aliases: []string{"b62"},
			Usage:   "Generate a base62 string",
			Action:  Base62,
			Flags: []cli.Flag{
				&lengthFlag,
				&prefixFlag,
			},
		},
		{
			Name:    "number",
			Aliases: []string{"n"},
			Usage:   "Generate a random number",
			Action:  Number,
			Flags: []cli.Flag{
				&lengthFlag,
			},
		},
		{
			Name:    "hex",
			Aliases: []string{""},
			Usage:   "Generate a random hex string",
			Action:  Hex,
			Flags: []cli.Flag{
				&bitFlag,
				&byteFlag,
			},
		},
		{
			Name:    "password",
			Aliases: []string{"pass"},
			Usage:   "Generate a random password",
			Action:  Password,
			Flags: []cli.Flag{
				&lengthFlag,
			},
		},
		{
			Name:    "base32",
			Aliases: []string{"b32"},
			Usage:   "Generate a base32 string",
			Action:  Base32,
			Flags: []cli.Flag{
				&lengthFlag,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
