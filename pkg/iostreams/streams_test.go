package iostreams_test

import (
	"os"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"

	"github.com/alex-held/dfctl-kit/pkg/iostreams"
	"github.com/alex-held/dfctl-kit/pkg/testutils"
	. "github.com/alex-held/dfctl-kit/pkg/testutils/matchers"
)

func TestDefault(t *testing.T) {
	testutils.Run(t, "Default()", func(g *goblin.G) {
		var sut = iostreams.Default()

		g.It("Set Out to Stdout", func() {
			Ω(sut.Out).Should(BeFile(os.Stdout))
		})

		g.It("Set Err to Stderr", func() {
			Ω(sut.Err).Should(BeFile(os.Stderr))
		})

		g.It("Set In to Stdin", func() {
			testutils.ConfigureTestLogger(t)
			log.Debug().Msgf("my debug message")
			Ω(sut.In).Should(BeFile(os.Stdin))
		})
	})
}
