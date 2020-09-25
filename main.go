package main

import (
	"strings"
	"github.com/pkg/errors"
	"os"
	"fmt"
)

func main() {
	if err := echo(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+vn", err)
	}
	os.Exit(1)
}

func echo(args []string) error {
	if len(args) < 2 {
		return errors.New("No message to echo")
	}
	_, err := fmt.Println(strings.Join(args[1:], " "))
	return err
}