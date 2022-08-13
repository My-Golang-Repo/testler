package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("ERROR", err)
		}
		fmt.Println(line)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("istifadə qaydasi: byLine <fayl1> <file2>\n")
		return
	}
	for _, file := range flag.Args() {
		err := lineByLine(file)

		if err != nil {
			fmt.Println(err)
		}
	}

}
