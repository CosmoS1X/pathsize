package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"

	"code"
)

func run(args []string, stdout, stderr io.Writer) int {
	cmd := &cli.Command{
		Name:                   "hexlet-path-size",
		Usage:                  "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage:              "<path>",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			//nolint:goconst // help output intentionally repeats the same default text
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Value:       false,
				Usage:       "recursive size of directories",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Value:       false,
				Usage:       "human-readable sizes (auto-select unit)",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Value:       false,
				Usage:       "include hidden files and directories",
				DefaultText: "false",
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
			all := c.Bool("all")
			recursive := c.Bool("recursive")

			size, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}
			fmt.Fprintf(stdout, "%s\t%s\n", size, path)
			return nil
		},
	}

	err := cmd.Run(context.Background(), args)
	if err != nil {
		fmt.Fprintf(stderr, "Error: %s\n", err.Error())
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}
