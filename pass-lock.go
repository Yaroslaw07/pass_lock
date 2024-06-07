package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	add := flag.Bool("add", false, "Add a new entry")

	flag.Parse()
	fmt.Printf("Password:")

	password, err := term.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	txt, err := readEnc(string(password))

	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}

	if *add {
		fmt.Printf("\nNew entry:")
		reader := bufio.NewReader(os.Stdin)
		entry, _ := reader.ReadString('\n')

		txt += entry
		err = writeEnc(txt, string(password))

		if err != nil {
			panic(err)
		}

		return
	}

	lines := strings.Split(strings.TrimSuffix(txt, "\n"), "\n")
	runUI(lines)
}
