package main

import (
	"errors"
	"fmt"

	generator "github.com/hacdan/rnd/internal"
	"github.com/urfave/cli/v2"
)

func hex(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("No arguments present") //TODO: better error messages
	}

	length := c.Int("length")

	s, err := generator.GenerateRandomHex(length)
	if err != nil {
		return err
	}

	fmt.Println(s) // TODO: Change to different output writer?
	return nil
}
