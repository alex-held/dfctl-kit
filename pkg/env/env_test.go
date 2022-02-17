package env

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"

	"github.com/alex-held/dfctl-kit/pkg/testutils"
)

func TestOverrideDefaults(t *testing.T) {
	sut := DefaultConfig(Vars{})

	sut.override(&Config{Home: "/foo"})

	assert.Equal(t, "/foo", sut.Home)
}

func TestLoad(t *testing.T) {

	testutils.Run(t, "Load", func(g *goblin.G) {

		g.BeforeEach(func() {
			ClearOverrides()
		})

		g.Describe("defaults", func() {
			g.It("sets Home to os.GetEnv(HOME)/.config/dfctl", func() {
				Ω(Load().Home).Should(Equal(os.Getenv("HOME") + "/.config/dfctl"))
			})

			g.It("sets ConfigFile to os.GetEnv(HOME)/.config/dfctl/dfctl.yaml", func() {
				Ω(Load().ConfigFile).Should(Equal(os.Getenv("HOME") + "/.config/dfctl/dfctl.yaml"))
			})
		})

		g.Describe("with config overrides and env overrides", func() {
			g.It("set overrides values", func() {
				Overrides.Home = "/override/dfctl"
				Overrides.ConfigFile = "/override/dfctl.yaml"
				Overrides.Vars = Vars{"DFCTL_CONFIG": "/home/foo/dfctl.yaml", "DFCTL_HOME": "/home/foo", "XDG_CONFIG_HOME": "/home/baz", "HOME": "/home/bar"}

				Ω(Load().Home).Should(Equal("/override/dfctl"))
				Ω(Load().ConfigFile).Should(Equal("/override/dfctl.yaml"))
			})
		})

		g.Describe("with config overrides", func() {

			g.JustBeforeEach(func() {
				ClearOverrides()
			})

			g.Describe("with Home override", func() {
				g.JustBeforeEach(func() {
					ClearOverrides()
				})

				g.It("sets Home to '/override/home", func() {
					Overrides.Home = "/override/home"
					Ω(Load().Home).Should(Equal("/override/home"))
				})
				g.It("sets ConfigFile to 'os.GetEnv($HOME)/.config/dfctl/dfctl.yaml", func() {
					Overrides.Home = "/override/home"
					Ω(Load().ConfigFile).Should(Equal(filepath.Join(os.Getenv("HOME"), ".config", "dfctl", "dfctl.yaml")))
				})
			})

			g.Describe("with ConfigFile override", func() {
				g.It("sets Home to 'os.GetEnv(HOME)/.config/dfctl", func() {
					Overrides.ConfigFile = "/override/dfctl.yaml"
					Ω(Load().Home).Should(Equal(filepath.Join(os.Getenv("HOME"), ".config", "dfctl")))
				})
				g.It("sets ConfigFile to 'os.GetEnv($HOME)/.config/dfctl/dfctl.yaml", func() {
					Overrides.ConfigFile = "/override/dfctl.yaml"
					Ω(Load().ConfigFile).Should(Equal("/override/dfctl.yaml"))
				})
			})
		})

		g.Describe("with env overrides", func() {

			g.Describe("with DFCTL_HOME, XDG_CONFIG_HOME, HOME variables", func() {
				g.JustBeforeEach(func() {
					Overrides.Vars = Vars{"DFCTL_HOME": "/home/foo", "XDG_CONFIG_HOME": "/home/baz", "HOME": "/home/bar"}
				})

				g.It("sets Home to 'DFCTL_HOME'", func() {
					Ω(Load().Home).Should(Equal("/home/foo"))
				})

				g.It("sets ConfigFile to '/home/foo/dfctl.yaml'", func() {
					Ω(Load().ConfigFile).Should(Equal("/home/foo/dfctl.yaml"))
				})
			})

			g.Describe("with DFCTL_HOME variable", func() {
				g.It("sets Home to '/home/foo'", func() {
					Ω(DefaultConfig(Vars{"DFCTL_HOME": "/home/foo"}).Home).Should(Equal("/home/foo"))
				})

				g.It("sets ConfigFile to '/home/foo/dfctl.yaml'", func() {
					Ω(DefaultConfig(Vars{"DFCTL_HOME": "/home/foo"}).ConfigFile).Should(Equal("/home/foo/dfctl.yaml"))
				})
			})

			g.Describe("with XDG_CONFIG_HOME variable", func() {
				g.It("sets Home to '/home/baz/dfctl'", func() {
					Ω(DefaultConfig(Vars{"XDG_CONFIG_HOME": "/home/baz"}).Home).Should(Equal("/home/baz/dfctl"))
				})

				g.It("sets ConfigFile to '/home/baz/dfctl/dfctl.yaml'", func() {
					Ω(DefaultConfig(Vars{"XDG_CONFIG_HOME": "/home/baz"}).ConfigFile).Should(Equal("/home/baz/dfctl/dfctl.yaml"))
				})
			})

			g.Describe("with HOME variable", func() {
				g.It("sets Home to '/home/bar/.config/dfctl'", func() {
					Ω(DefaultConfig(Vars{"HOME": "/home/bar"}).Home).Should(Equal("/home/bar/.config/dfctl"))
				})

				g.It("sets ConfigFile to '/home/bar/.config/dfctl/dfctl.yaml'", func() {
					Ω(DefaultConfig(Vars{"HOME": "/home/bar"}).ConfigFile).Should(Equal("/home/bar/.config/dfctl/dfctl.yaml"))
				})
			})

			g.Describe("without any environment variables", func() {
				g.It("sets Home to '.config/dfctl'", func() {
					Ω(DefaultConfig(Vars{}).Home).Should(Equal(".config/dfctl"))
				})

				g.It("sets ConfigFile to '.config/dfctl/dfctl.yaml'", func() {
					Ω(DefaultConfig(Vars{}).ConfigFile).Should(Equal(".config/dfctl/dfctl.yaml"))
				})
			})
		})
	})

	// tt := []struct {
	// 	name      string
	// 	env       map[string]string
	// 	want      Environment
	// 	overrides Config
	// }{
	// 	{
	// 		name: "with environment values",
	// 		env: Vars{
	// 			"DFCTL_HOME": "/env/home",
	// 		},
	// 		want: Environment{
	// 			Config: Config{
	// 				Home: "/env/home",
	// 			},
	// 			Vars: Vars{
	// 				"DFCTL_HOME": "/env/home",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		name: "with environment values and Overrides",
	// 		env: map[string]string{
	// 			"DFCTL_HOME":   "/env/home",
	// 			"DFCTL_CONFIG": "/env/home/dfctl.yaml",
	// 		},
	// 		want: Environment{
	// 			Config: Config{
	// 				Home:       "/override/home",
	// 				ConfigFile: "/env/home/dfctl.yaml",
	// 			},
	// 			Vars: map[string]string{
	// 				"DFCTL_HOME":   "/env/home",
	// 				"DFCTL_CONFIG": "/env/home/dfctl.yaml",
	// 			},
	// 		},
	// 		overrides: Config{
	// 			Home: "/override/home",
	// 		},
	// 	},
	// 	{
	// 		name: "defaults",
	// 		env: Vars{
	// 			"HOME": "$HOME",
	// 		},
	// 		want: Environment{
	// 			Vars: Vars{
	// 				"HOME": "$HOME",
	// 			},
	// 			Config: Config{
	// 				Home:       "$HOME/.config/dfctl",
	// 				ConfigFile: "$HOME/.config/dfctl/dfctl.yaml",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		name: "defaults with Overrides",
	// 		env: Vars{
	// 			"HOME": "$HOME",
	// 		},
	// 		want: Environment{
	// 			Vars: Vars{
	// 				"HOME": "$HOME",
	// 			},
	// 			Config: Config{
	// 				Home:       "/override/home",
	// 				ConfigFile: "/dfctl.yaml",
	// 			},
	// 		},
	// 		overrides: Config{
	// 			Home: "/override/home",
	// 		},
	// 	},
	// }
	//
	// for _, tt := range tt {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		Overrides.Vars = tt.env
	// 		Overrides.Config = tt.overrides
	//
	// 		got := Load()
	//
	// 		assert.Equal(t, tt.want, got)
	// 	})
	// }
}

func TestOsEnvironment(t *testing.T) {
	want := len(os.Environ())
	got := len(OsVars())
	assert.Equal(t, want, got)
}
