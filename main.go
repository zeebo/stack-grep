package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	invert = flag.Bool("v", false,
		"if true, will print things that don't match")
)

func main() {
	flag.Parse()
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() (err error) {
	input := strings.Join(flag.Args(), " ")

	re, err := regexp.Compile(input)
	if err != nil {
		return err
	}

	var stack []string
	lines := bufio.NewReader(os.Stdin)
	for {
		line, err := lines.ReadString('\n')
		if err == io.EOF {
			match(re, stack)
			return nil
		}
		if err != nil {
			return err
		}

		if line != "\n" {
			stack = append(stack, line)
			continue
		}

		match(re, stack)
		stack = stack[:0]
	}
}

func match(re *regexp.Regexp, stack []string) {
	if full := strings.Join(stack, ""); re.MatchString(full) != *invert {
		fmt.Println(full)
	}
}
