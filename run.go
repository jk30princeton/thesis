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

	command3 := exec.Command("/bin/bash", "-c", "nix-store -qR $(nix-instantiate '<nixpkgs>' -A firefox)")
	out3 := strings.TrimSpace(run(command3))
	fmt.Println(out3)

	split := strings.Split(out3, "\n")

	for i, s := range split {
		if i == 1 {
			fmt.Println("finished")
			break
		}
		if s == "" {
			continue
		}
		fmt.Println(i, s)
	}

	// nixStore := getNixStore()

	// var derivation [1]string
	// derivation[0] =

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
