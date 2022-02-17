package env

import (
	"path/filepath"
)

func Home() string       { return Load().Home }
func ConfigFile() string { return Load().ConfigFile }

func OMZ() string     { return filepath.Join(Home(), "omz") }
func Themes() string  { return filepath.Join(OMZ(), "custom", "themes") }
func Plugins() string { return filepath.Join(OMZ(), "custom", "plugins") }

func Extensions() string { return filepath.Join(Home(), "extensions") }
func Configs() string    { return filepath.Join(Home(), "configs") }
func Data() string       { return filepath.Join(Home(), "data") }
func SDKs() string       { return filepath.Join(Data(), "sdk") }
