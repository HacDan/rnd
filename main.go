package main

import (
	"flag"
	"fmt"
	"os"

	generator "github.com/hacdan/rnd/internal"
)

var base62Flag = flag.Int("62", 16, "rnd -62 <length>")                          // TODO: Revisit default value
var hexFlag = flag.Int("hex", 16, "rnd -hex <length in bytes>")                  // TODO: Revisit default value
var numFlag = flag.Int("num", 10, "rnd -num <length>")                           // TODO: Revisit default value
var prefixFlag = flag.String("prefix", "", "rnd -62 <length> --prefix \"gho_\"") // TODO: Revisit default value

func main() {
	flag.Parse()
	//s, _ := generator.GenerateRandomHex(128)
	if base62Flag == nil {
		fmt.Println("Empty flag")
		os.Exit(1)
	}

	s, err := generator.GenerateRandomStringBase62(*base62Flag, "")
	fmt.Println(*base62Flag)
	if err != nil {
		panic(err) //TODO: Don't panic, exit gracefully with reasoning
	}
	fmt.Println(s)

}
