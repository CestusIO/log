package log

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

// A Level is a logging priority. Lower levels are more important.
type Level int

const (
	// Error is used to log golang errors with stacktrace
	Error Level = iota - 2
	// Warning is a warning that something is wrong but it is mitigated.
	Warning
	// Info  is something you always want to show.
	Info
	//Debug1 is a debug level which is verbose. This should normally not be enabled in production
	Debug1
	//Debug2 is even more verbose then Debug1
	Debug2
	//Debug3 is even more verbose then Debug2
	Debug3
	//Debug4 is even more verbose then Debug3
	Debug4
	//Debug5 is even more verbose then Debug4
	Debug5
	//Debug6 is even more verbose then Debug5
	Debug6
	//Debug7 is even more verbose then Debug6
	Debug7
	//Debug8 is even more verbose then Debug7
	Debug8
	//Debug9 is even more verbose then Debug8
	Debug9
	//Debug10 is even more verbose then Debug9
	Debug10
)

// String returns a lower-case ASCII representation of the log level.
func (l Level) String() string {
	switch l {
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Info:
		return "info"
	case Debug1:
		return "debug1"
	case Debug2:
		return "debug2"
	case Debug3:
		return "debug3"
	case Debug4:
		return "debug4"
	case Debug5:
		return "debug5"
	case Debug6:
		return "debug6"
	case Debug7:
		return "debug7"
	case Debug8:
		return "debug8"
	case Debug9:
		return "debug9"
	case Debug10:
		return "debug10"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

// AsInt returns the level as an integer e.g. for the use in logr
func (l Level) AsInt() int {
	return int(l)
}

// CapitalString returns an all-caps ASCII representation of the log level.
func (l Level) CapitalString() string {
	// Printing levels in all-caps is common enough that we should export this
	// functionality.
	switch l {
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Info:
		return "INFO"
	case Debug1:
		return "DEBUG1"
	case Debug2:
		return "DEBUG2"
	case Debug3:
		return "DEBUG3"
	case Debug4:
		return "DEBUG4"
	case Debug5:
		return "DEBUG5"
	case Debug6:
		return "DEBUG6"
	case Debug7:
		return "DEBUG7"
	case Debug8:
		return "DEBUG8"
	case Debug9:
		return "DEBUG9"
	case Debug10:
		return "DEBUG10"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

// MarshalText marshals the Level to text.
func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalText unmarshals text to a level.
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "error", "ERROR":
		*l = Error
	case "warning", "WARNING":
		*l = Warning
	case "info", "INFO":
		*l = Info
	case "debug1", "DEBUG1":
		*l = Debug1
	case "debug2", "DEBUG2":
		*l = Debug2
	case "debug3", "DEBUG3":
		*l = Debug3
	case "debug4", "DEBUG4":
		*l = Debug4
	case "debug5", "DEBUG5":
		*l = Debug5
	case "debug6", "DEBUG6":
		*l = Debug6
	case "debug7", "DEBUG7":
		*l = Debug7
	case "debug8", "DEBUG8":
		*l = Debug8
	case "debug9", "DEBUG9":
		*l = Debug9
	case "debug10", "DEBUG10":
		*l = Debug10
	default:
		return false
	}
	return true
}

// VersionString is the version number
type VersionString string

// NewZapDevelopmentConfig is a development logger config
func NewZapDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(-3),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// SetLevel sets the logLevel
func SetLevel(level Level, config *zap.Config) {
	config.Level.SetLevel(zapcore.Level(-1 * level))
}

// ProvideZaprNoOpLogger provides a null zapr logger
func ProvideZaprNoOpLogger(version VersionString) (logr.Logger, error) {
	zapLog := zap.NewNop()

	return zapr.NewLogger(zapLog), nil
}

// ProvideZaprLogger provides a logger
func ProvideZaprLogger(version VersionString, config zap.Config) (logr.Logger, error) {
	var log logr.Logger

	zapLog, err := config.Build()
	if err != nil {
		zapLog := zap.NewNop()
		return zapr.NewLogger(zapLog), err
	}
	log = zapr.NewLogger(zapLog).WithValues("version", version)

	return log, nil
}
