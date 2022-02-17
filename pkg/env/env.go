package env

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var Overrides = &Environment{
	Vars:   Vars{},
	Config: Config{},
}

type Vars map[string]string

func (v Vars) WithoutOs() (vars Vars) {
	osVars := OsVars()
	for key, osVal := range osVars {
		if val, ok := v[key]; ok {
			if val != osVal {
				vars[key] = val
			}
		}
	}
	return vars
}

func (v Vars) override(vars Vars) {
	for key, val := range vars {
		v[key] = val
	}
}

func (v Vars) Lookup(name string) (val string, ok bool) {
	val, ok = v[name]
	return val, ok
}

func (v Vars) GetOrDefault(name string, defaults string) string {
	if val, ok := v[name]; ok {
		return val
	}
	return defaults
}

func (v Vars) Get(name string) string {
	return v.GetOrDefault(name, "")
}

func (v Vars) IsSet(name string) bool {
	_, ok := v.Lookup(name)
	return ok
}

func (v Vars) List() (list []string) {
	for key, val := range v {
		list = append(list, fmt.Sprintf("%s=%s", key, val))
	}
	return list
}
func (v Vars) With(vars Vars) Vars {
	for key, val := range vars {
		v[key] = val
	}
	return v
}

func (v Vars) HasAny() bool {
	return len(v) > 0
}

type EnvironmentOption func(cfg Vars)

func OsVars() (env Vars) {
	env = Vars{}
	for _, e := range os.Environ() {
		i := strings.Index(e, "=")
		envKey := e[:i]
		envVal := e[i+1:]
		env[envKey] = envVal
	}
	return env
}

func ClearOverrides() {
	Overrides.Config = Config{}
	Overrides.Vars = Vars{}
}

type Config struct {
	Home       string `env:"DFCTL_HOME,default=$HOME/.config/dfctl"`
	ConfigFile string `env:"DFCTL_CONFIG,default=$HOME/.config/dfctl/dfctl.yaml"`
}

func (e *Config) override(cfg *Config) {
	oVal := reflect.ValueOf(cfg).Elem()
	numI := oVal.NumField()
	for i := 0; i < numI; i++ {
		field := oVal.Field(i)
		if !field.IsZero() {
			reflect.ValueOf(e).Elem().Field(i).Set(field)
		}
	}
}

func Load() Environment {
	vars := OsVars()
	vars.override(Overrides.Vars)

	cfg := DefaultConfig(vars)
	cfg.override(&Overrides.Config)

	return Environment{
		Vars:   vars,
		Config: cfg,
	}
}

func DefaultConfig(vars Vars) (cfg Config) {
	var home, configFile string

	switch {
	case vars.IsSet("DFCTL_HOME"):
		home = filepath.Join(vars.Get("DFCTL_HOME"))
	case vars.IsSet("XDG_CONFIG_HOME"):
		home = filepath.Join(vars.Get("XDG_CONFIG_HOME"), "dfctl")
	case vars.IsSet("HOME"):
		home = filepath.Join(vars.Get("HOME"), ".config", "dfctl")
	default:
		home = ".config/dfctl"
	}

	configFile = filepath.Join(home, "dfctl.yaml")

	if vars.IsSet("DFCTL_CONFIG") {
		configFile = vars.Get("DFCTL_CONFIG")
	}

	return Config{
		Home:       home,
		ConfigFile: configFile,
	}
}

func GetVars() (vars Vars) {
	return OsVars().With(Overrides.Vars)
}

type Environment struct {
	Vars
	Config
}
