package jotnar

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type config interface {
	witchConfig()
}

// if isCommandConfig true, then your application should
// run as follow
// ./application --key1 value1 \
// --key2 value2 \
// --key3 value3
// if isCommandConfig false, use config file to init like follow
// ./application -f configfile
type defaultConfig struct {
	IsCommandConfig bool
	Config          map[string]string // config key value
}

type viperConfig struct {
	v *viper.Viper
}

var ConfigFileType = struct {
	Json       string
	Toml       string
	Yaml       string
	Hcl        string
	Ini        string
	Properties string
}{"json", "toml", "yaml", "hcl", "ini", "properties"}

func (*defaultConfig) witchConfig() {}
func (*viperConfig) witchConfig()   {}

// when use viper, set value 'viper'
var CurrentConfigType = "default"

func NewDefaultConfig() *defaultConfig {
	return &defaultConfig{
		IsCommandConfig: true,
		Config:          make(map[string]string),
	}
}

// use -f to appoint a config file
func NewViperConfig(fileType string) *viperConfig {
	if len(os.Args) < 3 {
		panic("need a config file use -f")
	}

	if os.Args[1] != "-f" {
		panic("must use -f")
	}

	v := viper.New()
	v.SetConfigFile(os.Args[2])
	v.SetConfigType(fileType)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return &viperConfig{v}
}

func NewViperConfigToml() *viperConfig {
	return NewViperConfig(ConfigFileType.Toml)
}

func NewViperConfigTest(fileType, filePath string) *viperConfig {
	v := viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType(fileType)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return &viperConfig{v}
}

func NewViperConfigTomlTest(filePath string) *viperConfig {
	return NewViperConfigTest(ConfigFileType.Toml, filePath)
}

var (
	ViperConfig   *viperConfig
	DefaultConfig *defaultConfig
)

// after init, the config will load in memory
// then you guys can use by GetValue...
func (j *Jotnar) InitConfig(c config) *Jotnar {
	switch cf := c.(type) {
	case *defaultConfig:
		CurrentConfigType = "default"
		DefaultConfig = cf
		doDefaultConfig(cf)
	case *viperConfig:
		CurrentConfigType = "viper"
		ViperConfig = cf
	default:
		panic("should use defaultConfig or viper to init")
	}

	return j
}

// use default k v flag by command
func (j *Jotnar) InitConfigDefaultCommandFlag() *Jotnar {
	return j.InitConfig(NewDefaultConfig())
}

// use viper to manage config; config file type is toml
func (j *Jotnar) InitConfigViperToml() *Jotnar {
	return j.InitConfig(NewViperConfigToml())
}

// for unit test, use this one
func (j *Jotnar) InitConfigViperTomlTest(filePath string) *Jotnar {
	return j.InitConfig(NewViperConfigTomlTest(filePath))
}

func doDefaultConfig(cf *defaultConfig) {
	if cf.IsCommandConfig == true {
		if len(os.Args) < 2 {
			panic("there is no arg, please input args or do not use this.")
		}
		args := os.Args[1:]

		if len(args)%2 != 0 {
			panic("some arg dont have value, please input it.")
		}

		for i := 0; i < len(args); i += 2 {
			cf.Config[strings.Trim(args[i], "-")] = strings.Trim(args[i+1], " ")
		}
	} else {
		// todo:
	}
}

// when use the DefaultConfig, you can use this to get value
// if dont have the key, then return <nil>
func GetValue(key string) string {
	if v, ok := DefaultConfig.Config[key]; !ok {
		return "<nil>"
	} else {
		return v
	}
}

// if you use viper, use this function
func GetViper() *viper.Viper {
	return ViperConfig.v
}

func GetString(key, defaultValue string) string {
	value := GetViper().GetString(key)
	if value == "" {
		return defaultValue
	}
	return value
}
