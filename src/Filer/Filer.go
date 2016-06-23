package main

import ( 
	"flag"
	"fmt" 
	"io/ioutil" 
	"log" 
)

var (
	format = flag.String("format", "", "the format you want your file in. P gives pointer, s=string, h=hex,g=go syntax, t=type")
	path = flag.String("path", "", "the path to the file you want to fetch data from.")
	filename = flag.String("filename", "", "the path to the file you want to violently assrape")
)

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(*path)
		if err != nil { 
			log.Fatal(err) 
		}
	switch { 
		case *format == "p":
			fmt.Printf("Location: %p\n", &file[0])
		case *format == "s": 
			fmt.Printf("Data:\n%s", file) 
		case *format == "h":
			fmt.Printf("Data: %x\n", file) 
		case *format == "g":
			fmt.Printf("Go syntax: %#v\n", file) 
		case *format == "t":
			fmt.Printf("Type: %T\n", file)
	} 

	if *filename != "" {
		Err := ioutil.WriteFile(*filename, file, 0644)
		if Err != nil{
			log.Fatal(Err)
		}
	}
} 