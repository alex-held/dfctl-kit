package dfpath

import (
	"path/filepath"

	"github.com/alex-held/dfctl-kit/pkg/env"
)

func Home() string       { return env.MustLoad().Home }
func OMZ() string        { return env.MustLoad().OMZ }
func ConfigFile() string { return filepath.Join(Home(), "dfctl"+env.MustLoad().ConfigFileType) }
func Themes() string     { return filepath.Join(OMZ(), "custom", "themes") }
func Plugins() string    { return filepath.Join(OMZ(), "custom", "plugins") }
func Extensions() string { return filepath.Join(Home(), "extensions") }
func Configs() string    { return filepath.Join(Home(), "configs") }
func Data() string       { return filepath.Join(Home(), "data") }
func SDKs() string       { return filepath.Join(Data(), "sdk") }
