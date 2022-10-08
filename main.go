package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	structs "pradhamk/ninite4linux/structs"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var platform string
var StringFormatGreen string = color.GreenString("[*]")
var StringFormatRed string = color.RedString("[*]")

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

func GetOSPlatform() string {
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

func InstallPackage(pkg structs.Package) {
	success := true
	var cmds []string
	for i := 0; i < len(pkg.Platforms); i++ {
		if pkg.Platforms[i].Name == platform {
			cmds = pkg.Platforms[i].Commands
			break
		}
	}
	fmt.Printf("%s Installing %s\n", StringFormatGreen, pkg.Name)
	for i := 0; i < len(cmds); i++ {
		err := exec.Command("bash", "-c", cmds[i]).Run()
		if err != nil {
			success = false
			break
		}
	}
	if !success {
		fmt.Printf("%s Unable to install %s\n", StringFormatRed, pkg.Name)
	} else {
		fmt.Printf("%s Installed %s\n", StringFormatGreen, pkg.Name)
	}
}

func PromptUser(pkgs structs.Packages) []string {
	var PkgNames []string
	var SelectedPkgs []string
	for i := 0; i < len(pkgs.Packages); i++ {
		PkgNames = append(PkgNames, pkgs.Packages[i].Name)
	}
	PkgNames = append(PkgNames, "Install?")
	for {
		contains := false
		prompt := promptui.Select{
			Label: "Select the packages you want to install:",
			Items: PkgNames,
		}

		_, result, _ := prompt.Run()
		if result == "Install?" {
			break
		} else if len(result) == 0 {
			continue
		}

		for _, val := range SelectedPkgs {
			if result == val {
				contains = true
				break
			}
		}
		if !contains {
			SelectedPkgs = append(SelectedPkgs, result)
		}
	}
	return SelectedPkgs
}

func main() {
	packages := ReadPackagesList()
	platform = GetOSPlatform()
	//platform = "debian"
	fmt.Printf("%s OS Platform detected as %s\n", StringFormatGreen, strings.ToUpper(platform))
	SelectedPkgs := PromptUser(packages)
	for i := 0; i < len(SelectedPkgs); i++ {
		var pkg structs.Package
		for j, val := range packages.Packages {
			if val.Name == SelectedPkgs[i] {
				pkg = packages.Packages[j]
				break
			}
		}
		InstallPackage(pkg)
	}
}
