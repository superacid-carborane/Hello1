//second timer

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func looper(num int64, i int64) string {
	switch {
	case i == num:
		return "Time's up! \n"
	case i != num:
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i + 1)
		i += 1
		looper(num, i)
	}
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter time (seconds): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	number, err := strconv.ParseInt(strings.TrimSuffix(text, "\n"), 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(looper(number, 0))
}
