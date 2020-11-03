package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Info struct {
	InputDerivations map[string][]string `json:"inputDrvs"`
}

type Derivations map[string]Info

func main() {
	// var fileName string
	// flag.StringVar(&fileName, "f", "", "JSON file to parse.")
	// flag.Parse()

	// if fileName == "" {
	// 	fmt.Println("Please provide json file by using -f option")
	// 	return
	// }

	// jsonFile, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Printf("Error reading JSON file: %s\n", err)
	// 	return
	// }

	dir := "/Users/josephkim/Documents/Senior2020/deplorable"

	command1 := exec.Command("cp", "linearscript.sh", dir)
	_, _ = run(command1)

	command2 := exec.Command("./linearscript.sh")
	command2.Dir = dir
	out2, _ := run(command2)

	command3 := exec.Command("rm", "linearscript.sh")
	command3.Dir = dir
	_, _ = run(command3)

	nixcommand := "nix show-derivation " + out2
	fmt.Println(nixcommand)
	command := exec.Command("/bin/bash", "-c", nixcommand, dir)
	out, _ := run(command)
	fmt.Println()
	fmt.Println()
	fmt.Println(out)

	dec := json.NewDecoder(strings.NewReader((out)))
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
