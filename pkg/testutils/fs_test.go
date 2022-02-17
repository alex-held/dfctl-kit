package testutils_test

import (
	"path/filepath"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/alex-held/dfctl-kit/pkg/testutils"
	. "github.com/alex-held/dfctl-kit/pkg/testutils/matchers"
)

func TestTempDir(t *testing.T) {
	testutils.Run(t, "TempDir()", func(g *goblin.G) {

		g.It("returns path under the os.TempDir directory", func() {
			got := testutils.TempDir(t)
			Ω(got).Should(BeTempPath(afero.NewOsFs()))
		})

		g.It("returns path under the temp dfctl-test directory", func() {
			got := testutils.TempDir(t)
			Ω(got).Should(MatchRegexp(".*/dfctl-test/.*"))
		})

		g.It("returns path in the test-name directory", func() {
			got := testutils.TempDir(t)
			Ω(filepath.Base(got)).Should(BeNamedFileOrDir(t.Name()))
		})

		g.Describe("Without additional paths", func() {
			g.It("returns os tmp dir with `dfctl-test/<TEST-NAME>` suffix", func() {
				got := testutils.TempDir(t)
				Ω(filepath.Base(got)).Should(BeNamedFileOrDir(t.Name()))
			})
		})

		g.Describe("With additional paths", func() {
			g.It("returns path with additional path suffix", func() {
				got := testutils.TempDir(t, "foo")
				Ω(filepath.Base(got)).Should(Equal("foo"))
			})
			g.It("returns path under temp test dir", func() {
				got := testutils.TempDir(t, "foo")
				Ω(filepath.Base(filepath.Dir(got))).Should(Equal(t.Name()))
			})
		})
	})
}
