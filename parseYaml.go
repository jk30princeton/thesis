package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Repos struct {
	Repositories map[string][]Repo `yaml:"repos"`
}

type Repo map[string]string

// type Repos1 struct {
// 	Repositories []Repo1 `yaml:"repos"`
// }

// type Repo1 map[string]string

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "YAML file to parse.")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide yaml file by using -f option")
		return
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	repos := parseYaml(yamlFile)

	fmt.Println(repos)
}

func parseYaml(yamlFile []uint8) Repos {
	var repos Repos
	err := yaml.Unmarshal(yamlFile, &repos)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	return repos
}
