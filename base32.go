package main

import (
	"fmt"

	generator "github.com/hacdan/rnd/internal"
	"github.com/urfave/cli/v2"
)

func Base32(c *cli.Context) error {
	length := c.Int("length")

	s, err := generator.GenerateBase32(length)
	if err != nil {
		return err
	}

	fmt.Println(s) // TODO: Change to different output writer?
	return nil
}
