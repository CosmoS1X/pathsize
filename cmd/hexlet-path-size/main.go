package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"code"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.Args().Get(0)
			size, err := code.GetPathSize(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
				os.Exit(1)
			}

			fmt.Printf("%s\t%s\n", size, path)

			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
