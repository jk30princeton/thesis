package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// content, err := ioutil.ReadFile("/Users/josephkim/Downloads/names.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// text := string(content)
	// split := strings.Split(text, "\n")

	// command2 := exec.Command("/bin/bash", "-c", "nix-instantiate '<nixpkgs>' -A firefox")
	// out2 := strings.TrimSpace(run(command2))
	// fmt.Println(out2)
	// var derivation [1]string

	// for i, s := range split {
	// 	if i == 1 {
	// 		fmt.Println("finished")
	// 		break
	// 	}
	//
	// 	if s == "" {
	// 		continue
	// 	}
	//	derivation[0] = s
	//  sum := sum(derivation, )
	// 	fmt.Println(i, s)
	// }

	command := exec.Command("/bin/bash", "-c", "nix-instantiate '<nixpkgs>' -A firefox")
	out := strings.TrimSpace(run(command))
	fmt.Println(out)

	command3 := exec.Command("/bin/bash", "-c", "nix show-derivation -r "+out)
	out3 := strings.TrimSpace(run(command3))
	fmt.Println(out3)

	// dec := json.NewDecoder(strings.NewReader((out3)))
	// for {
	// 	var derivation Derivations
	// 	if err := dec.Decode(&derivation); err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("There was some error.")
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Println(derivation)

	// nixStore := getNixStore()

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
