package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
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

func getOSPlatform() string {
	out, err := exec.Command("cat", "/etc/os-release").Output()
	ErrHandle(err)
	fileData := string(out)
	data := strings.Split(fileData, "\n")
	var idId int
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i], "ID=") {
			idId = i
			break
		}
	}
	return strings.Replace(data[idId], "ID=", "", 1)
}

func main() {
	//packages := ReadPackagesList()
	//platform := getOSPlatform()
}
