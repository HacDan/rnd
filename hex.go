package main

import (
	"errors"
	"fmt"

	generator "github.com/hacdan/rnd/internal"
	"github.com/urfave/cli/v2"
)

func Hex(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("No arguments present") //TODO: better error messages
	}

	length := 0

	bits := c.Int("bits")
	bytes := c.Int("bytes")
	if bits == 0 && bytes > 0 {
		length = bytes
	}
	if bits > 0 && bytes == 0 {
		length = bits / 4 // Bits to bytes, man!
	}

	s, err := generator.GenerateRandomHex(length)
	if err != nil {
		return err
	}

	fmt.Println(s) // TODO: Change to different output writer?
	return nil
}
