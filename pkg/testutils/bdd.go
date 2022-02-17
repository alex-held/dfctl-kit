package testutils

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func Run(t *testing.T, name string, test func(g *goblin.G)) {
	g := goblin.Goblin(t)
	ConfigureTestLogger(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe(name, func() {
		test(g)
	})
}
