package jotnar

var ConfigFileType = struct {
	Json       string
	Toml       string
	Yaml       string
	Hcl        string
	Ini        string
	Properties string
}{"json", "toml", "yaml", "hcl", "ini", "properties"}
