package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Info struct {
	InputDerivations map[string][]string `json:"inputDrvs"`
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "JSON file to parse.")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide json file by using -f option")
		return
	}

	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	type Derivations map[string]Info

	dec := json.NewDecoder(strings.NewReader(string(jsonFile)))
	for {
		var derivation Derivations
		if err := dec.Decode(&derivation); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for k, v := range derivation {
			rootnode := newNode(k)
			recursiveAdd(rootnode, v.InputDerivations)
			// fmt.Printf("Derivation is %s.\n", k)

			// fmt.Println("Input derivations are:")
			// fmt.Println("%T", v.InputDerivations)
			// for key, value := range v.InputDerivations {
			// 	fmt.Printf("%T, %T", key, value)
			// 	fmt.Printf("%s\n", key)
			// }
		}
	}
}
