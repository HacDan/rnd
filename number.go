package main

import (
	"errors"
	"fmt"

	generator "github.com/hacdan/rnd/internal"
	"github.com/urfave/cli/v2"
)

func Number(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("No arguments present") //TODO: better error messages
	}

	length := c.Int("length")

	i, err := generator.GenerateRandomNum(length)
	if err != nil {
		return err
	}

	fmt.Println(i) // TODO: Change to different output writer?
	return nil
}
