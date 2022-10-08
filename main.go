package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	structs "pradhamk/ninite4linux/structs"
	"strings"
)

var platform string

func ErrHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ReadPackagesList() structs.Packages {
	var packages structs.Packages
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

func installPackage(pkg structs.Package) {
	var cmds []string
	for i := 0; i < len(pkg.Platforms); i++ {
		if pkg.Platforms[i].Name == platform {
			cmds = pkg.Platforms[i].Commands
			break
		}
	}
	for i := 0; i < len(cmds); i++ {
		s, err := exec.Command("bash", "-c", cmds[i]).Output()
		if err != nil {
			fmt.Printf("Unable to install package %s\n", pkg.Name)
		} else {
			fmt.Printf("%s successfully installed\n", pkg.Name)
		}
		fmt.Println(s)
	}
}

func main() {
	packages := ReadPackagesList()
	//platform = getOSPlatform()
	platform = "debian"
	installPackage(packages.Packages[0])
}
