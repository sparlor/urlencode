
dpackage main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	var encode bool
	var decode bool
	args := os.Args[2]

	flag.BoolVar(&encode, "e", false, "a bool")
	flag.BoolVar(&decode, "d", false, "a bool")
	flag.Parse()

	if decode {
		fmt.Println(url.PathUnescape(args))
	} else if encode {
		fmt.Println(url.PathEscape(args))
	} else {
		fmt.Println("Pick -encode or -decode")
	}
}
