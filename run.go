package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("/home/joseph/Downloads/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	split := strings.Split(text, "\n")

	nixStore := getNixStore()

	for i, s := range split {
		if i == 500 {
			fmt.Println("finished")
			break
		}

		if s == "" {
			continue
		}
		command := exec.Command("/bin/bash", "-c", "nix-instantiate '<nixpkgs>' -A "+s)
		out := strings.TrimSpace(run(command))
		fmt.Println(out)

		command3 := exec.Command("/bin/bash", "-c", "nix show-derivation -r "+out)
		out3 := strings.TrimSpace(run(command3))

		dec := json.NewDecoder(strings.NewReader((out3)))
		for {
			var derivation Derivations
			if err := dec.Decode(&derivation); err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("There was some error.")
				log.Fatal(err)
			}
			score := recursiveAdd(derivation[out].InputDerivations, derivation, 1, 1.0/float64(len(derivation[out].InputDerivations)), nixStore)
			fmt.Println("Score is ", score)
			fmt.Println(nixStore.Size())
		}
		fmt.Println(i, s)
		fmt.Println()
	}

	// set1 := strset.New()
	// set1.Add("Hello")

	// set2 := strset.New()
	// set2.Add("Hello")
	// set2.Add("World")

	// set3 := strset.Union(set1, set2)
	// set3.Add("Hello")
	// fmt.Println(set3)
	// fmt.Println(set3.Has("Hello"))
}
