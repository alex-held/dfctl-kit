package env_test

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"

	"github.com/alex-held/dfctl-kit/pkg/env"
	"github.com/alex-held/dfctl-kit/pkg/testutils"
	"github.com/alex-held/dfctl-kit/pkg/testutils/matchers"
)

func TestPaths(t *testing.T) {
	testutils.Run(t, "Paths", func(g *goblin.G) {

		defaultEnv := env.Config{
			Home:       "/home/foo/.config/dfctl",
			ConfigFile: "/home/foo/.config/dfctl/dfctl.yaml",
		}

		env.Overrides.Config = defaultEnv

		g.It("Home() should return Home", func() {
			Ω(env.Home()).Should(Equal(defaultEnv.Home))
		})

		g.It("OMZ() should return be under <home>/omz", func() {
			Ω(env.OMZ()).Should(matchers.BeSubPath(defaultEnv.Home, "omz"))
		})

		// func Home() string       { return env.MustLoad().Home }
		// func OMZ() string        { return env.MustLoad().OMZ }
		// func ConfigFile() string { return filepath.Join(Home(), "dfctl"+env.MustLoad().ConfigFileType) }
		// func Themes() string     { return filepath.Join(OMZ(), "custom", "themes") }
		// func Plugins() string    { return filepath.Join(OMZ(), "custom", "plugins") }
		// func Extensions() string { return filepath.Join(Home(), "extensions") }
		// func Configs() string    { return filepath.Join(Home(), "configs") }
		// func Data() string       { return filepath.Join(Home(), "data") }
		// func SDKs() string       { return filepath.Join(Data(), "sdk") }

	})
}
