package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Packages struct {
	Packages []Package `json: "packages"`
}

type Package struct {
	Name        string     `json: "name"`
	Description string     `json: "description"`
	Platforms   []Platform `json: "platforms"`
}

type Platform struct {
	Name     string   `json: "name"`
	Commands []string `json: "commands"`
}

func ErrHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ReadPackagesList() Packages {
	var packages Packages
	file, err := os.Open("packages.json")
	ErrHandle(err)
	defer file.Close()
	FileData, err := ioutil.ReadAll(file)
	ErrHandle(err)
	json.Unmarshal(FileData, &packages)
	return packages
}

func main() {
	packages := ReadPackagesList()
}
