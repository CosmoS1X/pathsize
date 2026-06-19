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
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "path",
			},
		},
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			argsLen := c.Args().Len()
			if argsLen == 0 {
				return ctx, fmt.Errorf("path argument is required")
			}
			if argsLen > 1 {
				return ctx, fmt.Errorf("too many arguments provided, expected only one path argument")
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.StringArg("path")
			human := c.Bool("human")

			size, err := code.GetPathSize(path, human)
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
