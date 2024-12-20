package config

import (
	_ "embed"
)

const (
	ChannelNightly = "nightly"
	ChannelRelease = "release"

	RootName   = ".flowx"
	BinName    = RootName + "/" + "bin"
	ConfigName = RootName + "/" + "flowup.yml"
	EnvName    = RootName + "/" + "env"

	BinPerm  = 0755
	DirPerm  = 0755
	FilePerm = 0644
)

var (
	Build   string
	Commit  string
	Version string
)

//go:embed config.yml
var ConfigData string

//go:embed env
var EnvData string

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type MetaData struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Artifact Artifact `yaml:"artifact"`
}

type Artifact struct {
	Url  string `yaml:"url"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}
