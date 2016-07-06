package filetools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Jot(path string, format string, path2 string, print bool) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	wfile, err := os.OpenFile(path2, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer wfile.Close()

	b := bufio.NewWriter(wfile)
	defer func() {
		if err = b.Flush(); err != nil {
			log.Fatal(err)
		}
	}()

	switch {
	case format == "p":
		fmt.Fprintf(b, "Location: %p\n", &file[0])
	case format == "s":
		fmt.Fprintf(b, "Data:\n%s", file)
	case format == "h":
		fmt.Fprintf(b, "Data: %x\n", file)
	case format == "g":
		fmt.Fprintf(b, "Go syntax: %#v\n", file)
	case format == "t":
		fmt.Fprintf(b, "Type: %T\n", file)
	}

	if print == true {
		switch {
		case format == "p":
			fmt.Printf("Location: %p\n", &file[0])
		case format == "s":
			fmt.Printf("Data:\n%s", file)
		case format == "h":
			fmt.Printf("Data: %x\n", file)
		case format == "g":
			fmt.Printf("Go syntax: %#v\n", file)
		case format == "t":
			fmt.Printf("Type: %T\n", file)
		}
	}
}
