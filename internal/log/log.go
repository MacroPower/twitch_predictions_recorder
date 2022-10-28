// Copyright 2017 The Prometheus Authors
// Modifications copyright 2021 Jacob Colvin (MacroPower)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package log defines standardized ways to initialize Go kit loggers.
// It should typically only ever be imported by main packages.
package log

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

const (
	logCallerDepth = 5
	jsonFormat     = "json"
	logfmtFormat   = "logfmt"
)

// This timestamp format differs from RFC3339Nano by using .000 instead
// of .999999999 which changes the timestamp from 9 variable to 3 fixed
// decimals (.130 instead of .130987456).
var timestampFormat = log.TimestampFormat(
	func() time.Time { return time.Now().UTC() },
	"2006-01-02T15:04:05.000Z07:00",
)

// AllowedLevel is a settable identifier for the minimum level a log entry
// must be have.
type AllowedLevel struct {
	s string
	o level.Option
}

func (l *AllowedLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	type plain string

	if err := unmarshal((*plain)(&s)); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	lo := &AllowedLevel{}
	if err := lo.Set(s); err != nil {
		return err
	}

	*l = *lo

	return nil
}

func (l *AllowedLevel) String() string {
	return l.s
}

// Set updates the value of the allowed level.
func (l *AllowedLevel) Set(s string) error {
	switch s {
	case "debug":
		l.o = level.AllowDebug()
	case "info":
		l.o = level.AllowInfo()
	case "warn":
		l.o = level.AllowWarn()
	case "error":
		l.o = level.AllowError()
	default:
		return fmt.Errorf("unrecognized log level %q", s)
	}

	l.s = s

	return nil
}

// AllowedFormat is a settable identifier for the output format that the logger can have.
type AllowedFormat struct {
	s string
}

func (f *AllowedFormat) String() string {
	return f.s
}

// Set updates the value of the allowed format.
func (f *AllowedFormat) Set(s string) error {
	switch s {
	case logfmtFormat, jsonFormat:
		f.s = s
	default:
		return fmt.Errorf("unrecognized log format %q", s)
	}

	return nil
}

// Config is a struct containing configurable settings for the logger.
type Config struct {
	Level  *AllowedLevel
	Format *AllowedFormat
}

// New returns a new leveled logger. Each logged line will be annotated
// with a timestamp. The output always goes to stderr.
func New(config *Config) log.Logger {
	var l log.Logger

	if config.Format != nil && config.Format.s == jsonFormat {
		l = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	} else {
		l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	}

	if config.Level != nil {
		l = log.With(l, "ts", timestampFormat, "caller", log.Caller(logCallerDepth))
		l = level.NewFilter(l, config.Level.o)
	} else {
		l = log.With(l, "ts", timestampFormat, "caller", log.DefaultCaller)
	}

	return l
}

// NewDynamic returns a new leveled logger. Each logged line will be annotated
// with a timestamp. The output always goes to stderr. Some properties can be
// changed, like the level.
func NewDynamic(config *Config) *Logger {
	var l log.Logger

	if config.Format != nil && config.Format.s == jsonFormat {
		l = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	} else {
		l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	}

	lo := &Logger{
		Base:    l,
		Leveled: l,
	}

	if config.Level != nil {
		lo.SetLevel(config.Level)
	}

	return lo
}

type Logger struct {
	Base         log.Logger
	Leveled      log.Logger
	CurrentLevel *AllowedLevel

	mtx sync.Mutex
}

// Log implements logger.Log.
func (l *Logger) Log(keyvals ...interface{}) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if err := l.Leveled.Log(keyvals...); err != nil {
		return fmt.Errorf("log error: %w", err)
	}

	return nil
}

// SetLevel changes the log level.
func (l *Logger) SetLevel(lvl *AllowedLevel) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if lvl == nil {
		l.Leveled = log.With(l.Base, "ts", timestampFormat, "caller", log.DefaultCaller)
		l.CurrentLevel = nil

		return
	}

	if l.CurrentLevel != nil && l.CurrentLevel.s != lvl.s {
		_ = l.Base.Log("msg", "Log level changed", "prev", l.CurrentLevel, "current", lvl)
	}

	l.CurrentLevel = lvl
	l.Leveled = level.NewFilter(log.With(l.Base, "ts", timestampFormat, "caller", log.Caller(logCallerDepth)), lvl.o)
}

// Error returns a logger that includes a Key/ErrorValue pair.
func Error(logger log.Logger) log.Logger {
	return log.WithPrefix(logger, level.Key(), level.ErrorValue())
}

// Warn returns a logger that includes a Key/WarnValue pair.
func Warn(logger log.Logger) log.Logger {
	return log.WithPrefix(logger, level.Key(), level.WarnValue())
}

// Info returns a logger that includes a Key/InfoValue pair.
func Info(logger log.Logger) log.Logger {
	return log.WithPrefix(logger, level.Key(), level.InfoValue())
}

// Debug returns a logger that includes a Key/DebugValue pair.
func Debug(logger log.Logger) log.Logger {
	return log.WithPrefix(logger, level.Key(), level.DebugValue())
}
