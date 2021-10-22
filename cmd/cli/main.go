package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/mertakman/deep-blue/src/pathgenerator"
)

var parseJson bool

func init() {
	flag.BoolVar(&parseJson, "parsejson", false, "Formats output into JSON")
	flag.Parse()
}

func main() {
	for {
		var src, dst string // variables for source and destination coordinates

		if n, err := fmt.Scanln(&src, &dst); err != nil {
			fmt.Println("Encountered with an error during the read of input : ", err)

			continue
		} else if n == 0 {
			// not read any input , returning an error.
			continue
		}

		if len(src) != 2 || len(dst) != 2 { // some hotpath check before actually sending the input .
			fmt.Println("Invalid input parameters , Coordinates should be 2 character , case insensitive strings. (ie: c3,h6,A1,B4 etc.)")

			continue
		}

		route, err := pathgenerator.GenerateKnightPath(src, dst)
		if err != nil {
			fmt.Println(err)

			continue
		}

		if parseJson {
			data, _ := json.Marshal(route)
			if err != nil {
				fmt.Println("Encountered with an error during marshaling an output : ", err.Error())
			}

			fmt.Printf("%s \n", data)
		} else {
			fmt.Printf("%s \n", strings.Join(route, " "))
		}
	}
}
