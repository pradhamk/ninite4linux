package structs

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
