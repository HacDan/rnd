package main

import (
	"fmt"

	generator "github.com/hacdan/rnd/internal"
	"github.com/urfave/cli/v2"
)

func Base62(c *cli.Context) error {
	length := c.Int("length")
	prefix := c.String("prefix")

	s, err := generator.GenerateRandomStringBase62(length, prefix)
	if err != nil {
		return err
	}

	fmt.Println(s) // TODO: Change to different output writer?
	return nil
}
