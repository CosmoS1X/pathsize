package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("Hello from Hexlet!")
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
