package testutils

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/alex-held/dfctl-kit/pkg/dflog"
)

func ConfigureTestLogger(t *testing.T) {
	dflog.ConfigureWithLevel(
		zerolog.TraceLevel,
		dflog.WithOrder(zerolog.LevelFieldName, "testcase", zerolog.MessageFieldName),
		dflog.WithFormatLevel(dflog.PrefixedFormatLevelFormatter(2, "➥")),
	)
	log.Logger = log.Logger.With().Str("testcase", t.Name()).Logger()
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
}
