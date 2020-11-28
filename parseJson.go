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
	dir := "/home/joseph/Documents/code/thesis"

	// command1 := exec.Command("cp", "linearscript.sh", dir)
	// run(command1)

	// command2 := exec.Command("./linearscript.sh")
	// command2.Dir = dir
	// out2 := strings.TrimSpace(run(command2))

	// fmt.Println(out2)

	// command3 := exec.Command("rm", "linearscript.sh")
	// command3.Dir = dir
	// run(command3)

	command2 := exec.Command("/bin/bash", "-c", "nix-instantiate default.nix")
	command2.Dir = dir
	out2 := strings.TrimSpace(run(command2))
	fmt.Println(out2)

	nixcommand := "nix show-derivation -r " + out2
	command := exec.Command("/bin/bash", "-c", nixcommand)
	command.Dir = dir
	out := run(command)
	fmt.Println()
	fmt.Println()
	fmt.Println(out)

	nixStore := getNixStore()

	dec := json.NewDecoder(strings.NewReader((out)))
	for {
		var derivation Derivations
		if err := dec.Decode(&derivation); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("hi")
			log.Fatal(err)
		}

		rootnode := newNode(out2)
		score := recursiveAdd(rootnode, derivation[out2].InputDerivations, derivation, 1, 1.0/float64(len(derivation[out2].InputDerivations)), nixStore)
		fmt.Println(rootnode)
		fmt.Println(score)
	}
}
