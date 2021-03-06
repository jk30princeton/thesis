package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/scylladb/go-set/strset"
)

type Info struct {
	InputDerivations map[string][]string `json:"inputDrvs"`
}

type Derivations map[string]Info

func main() {
	content, err := ioutil.ReadFile("/home/joseph/Downloads/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	split := strings.Split(text, "\n")
	rand.Seed(time.Now().UnixNano())
	for i := len(split) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		split[i], split[j] = split[j], split[i]
	}

	nixStore1 := strset.New()
	var nixStore1BuildList []string

	for i := 0; i < 3; i++ {
		command3 := exec.Command("/bin/bash", "-c", "nix-store -qR $(nix-instantiate '<nixpkgs>' -A "+split[i]+")")
		out3 := strings.TrimSpace(run(command3))
		split := strings.Split(out3, "\n")

		for j := range split {
			if split[j] == "" {
				continue
			}
			nixStore1.Add(split[j])
		}
	}
	nixStore2 := strset.New()
	var nixStore2BuildList []string
	for i := 3; i < 6; i++ {
		command3 := exec.Command("/bin/bash", "-c", "nix-store -qR $(nix-instantiate '<nixpkgs>' -A "+split[i]+")")
		out3 := strings.TrimSpace(run(command3))
		split := strings.Split(out3, "\n")

		for j := range split {
			if split[j] == "" {
				continue
			}
			nixStore2.Add(split[j])
		}
	}
	nixStore3 := strset.New()
	var nixStore3BuildList []string
	for i := 6; i < 9; i++ {
		command3 := exec.Command("/bin/bash", "-c", "nix-store -qR $(nix-instantiate '<nixpkgs>' -A "+split[i]+")")
		out3 := strings.TrimSpace(run(command3))
		split := strings.Split(out3, "\n")

		for j := range split {
			if split[j] == "" {
				continue
			}
			nixStore3.Add(split[j])
		}
	}

	// Build all the packages by assigning to the largest

	for i, s := range split {
		// if i == 17000 {
		// 	fmt.Println("finished")
		// 	break
		// }
		if i < 9 {
			continue
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
			score1 := recursiveAdd(derivation[out].InputDerivations, derivation, 1, 1.0/float64(len(derivation[out].InputDerivations)), nixStore1)
			score2 := recursiveAdd(derivation[out].InputDerivations, derivation, 1, 1.0/float64(len(derivation[out].InputDerivations)), nixStore2)
			score3 := recursiveAdd(derivation[out].InputDerivations, derivation, 1, 1.0/float64(len(derivation[out].InputDerivations)), nixStore3)
			fmt.Println(score1, score2, score3)
			builder := 0
			score := 0.0
			if score1 >= score2 && score1 >= score3 {
				score = score1
				builder = 1
			}
			if score2 >= score1 && score2 >= score3 {
				score = score2
				builder = 2
			}
			if score3 >= score1 && score3 >= score2 {
				score = score3
				builder = 3
			}

			fmt.Printf("Score is %f\n", score)
			fmt.Printf("Assigned to builder %d\n", builder)

			command3 := exec.Command("/bin/bash", "-c", "nix-store -qR "+out)
			out3 := strings.TrimSpace(run(command3))
			split := strings.Split(out3, "\n")
			if builder == 1 {
				nixStore1BuildList = append(nixStore1BuildList, out)
				for j := range split {
					if split[j] == "" {
						continue
					}
					nixStore1.Add(split[j])
				}
			} else if builder == 2 {
				nixStore2BuildList = append(nixStore2BuildList, out)
				for j := range split {
					if split[j] == "" {
						continue
					}
					nixStore2.Add(split[j])
				}
			} else if builder == 3 {
				nixStore3BuildList = append(nixStore3BuildList, out)
				for j := range split {
					if split[j] == "" {
						continue
					}
					nixStore3.Add(split[j])
				}
			}
		}
		fmt.Println(i, s)
		fmt.Println()
	}
	fmt.Println(nixStore1BuildList)
	fmt.Println(nixStore2BuildList)
	fmt.Println(nixStore3BuildList)
}
