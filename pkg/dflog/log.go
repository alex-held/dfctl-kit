package dflog

import (
	"fmt"
	"io"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/alex-held/dfctl-kit/pkg/color"
	"github.com/alex-held/dfctl-kit/pkg/iostreams"
)

type LoggerOption func(logger *zerolog.ConsoleWriter)

func WithFormatFieldValue(formatter zerolog.Formatter) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.FormatFieldValue = formatter
	}
}

func WithFormatFieldName(formatter zerolog.Formatter) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.FormatFieldName = formatter
	}
}

func DefaultFormatLevelFormatter() zerolog.Formatter {
	return PrefixedFormatLevelFormatter(0, "")
}

func PrefixedFormatLevelFormatter(indentation int, prefix string) zerolog.Formatter {
	prefixed := func(s string) string {
		if prefix == "" {
			return s
		}
		i := strings.Repeat(" ", indentation*4)
		return fmt.Sprintf("%s%s %s", i, color.Colorize(prefix, color.Green), s)
	}

	return func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case zerolog.LevelTraceValue:
				l = prefixed(color.Colorize("TRC", color.Magenta))
			case zerolog.LevelDebugValue:
				l = prefixed(color.Colorize("DBG", color.Cyan))
			case zerolog.LevelInfoValue:
				l = prefixed(color.Colorize("INF", color.Green))
			case zerolog.LevelWarnValue:
				l = prefixed(color.Colorize("WRN", color.Yellow))
			case zerolog.LevelErrorValue:
				l = prefixed(color.Colorize("ERR", color.Red, color.Bold))
			case zerolog.LevelFatalValue:
				l = prefixed(color.Colorize("FTL", color.Red, color.Bold))
			case zerolog.LevelPanicValue:
				l = prefixed(color.Colorize("PNC", color.Red, color.Bold))
			default:
				l = prefixed(color.Colorize("???", color.Bold))
			}
		} else {
			if i == nil {
				l = prefixed(color.Colorize("???", color.Bold))
			} else {
				l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
			}
		}
		return l
	}
}

func WithFormatLevel(formatter zerolog.Formatter) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.FormatLevel = formatter
	}
}

func WithOut(out io.Writer) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.Out = out
	}
}

func Without(parts ...string) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.PartsExclude = parts
	}
}

func WithOrder(parts ...string) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.PartsOrder = parts
	}
}

func WithColor(enabled bool) LoggerOption {
	return func(w *zerolog.ConsoleWriter) {
		w.NoColor = !enabled
	}
}

func Configure(opts ...LoggerOption) {
	ConfigureWithLevel(zerolog.InfoLevel, opts...)
}

func ConfigureWithLevel(level zerolog.Level, opts ...LoggerOption) {
	defaults := []LoggerOption{
		Without(zerolog.CallerFieldName, zerolog.TimestampFieldName),
		WithOrder(zerolog.LevelFieldName, zerolog.MessageFieldName),
		WithColor(true),
		WithOut(iostreams.Default().Out),
		WithFormatLevel(DefaultFormatLevelFormatter()),
	}

	w := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		for _, opt := range defaults {
			opt(w)
		}
		for _, opt := range opts {
			opt(w)
		}
	})

	log.Logger = zerolog.New(w)
	zerolog.SetGlobalLevel(level)
}

func ConfigureWithLevelString(levelString string, opts ...LoggerOption) {
	if level, err := zerolog.ParseLevel(levelString); err == nil {
		ConfigureWithLevel(level, opts...)
		return
	}
	Configure(opts...)
}
