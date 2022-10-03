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
	Name string `json: "name"`
}

func main() {
	file, err := os.Open("packages.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	FileData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	var packages Packages
	json.Unmarshal(FileData, &packages)
}
