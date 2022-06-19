package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	n := flag.Bool("n", false, "show line numbers")
	flag.Parse()

	files := flag.Args()
	path, err := os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stdout, "failed to read: ", err)
	}

	path = filepath.Dir(path)
	i := 1
	for j := 0; j < len(files); j++ {
		f, err := os.Open(filepath.Join(path, files[j]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to read: ", err)
		} else {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				if *n {
					fmt.Printf("%v: ", i)
				}
				fmt.Println(scanner.Text())
				i++
			}
		}
	}
}
