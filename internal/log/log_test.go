// Copyright 2020 The Prometheus Authors
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

package log_test

import (
	"fmt"
	"testing"

	"github.com/MacroPower/twitch_predictions_recorder/internal/log"
)

// Make sure creating and using a logger with an empty configuration doesn't
// result in a panic.
func TestDefaultConfig(t *testing.T) {
	t.Parallel()

	logger := log.New(&log.Config{})

	if err := logger.Log("hello", "world"); err != nil {
		t.Fatal(err)
	}
}

type recordKeyvalLogger struct {
	count int
}

func (r *recordKeyvalLogger) Log(keyvals ...interface{}) error {
	for _, v := range keyvals {
		if fmt.Sprintf("%v", v) == "Log level changed" {
			return nil
		}
	}
	r.count++

	return nil
}

func TestDynamic(t *testing.T) {
	t.Parallel()

	logger := log.NewDynamic(&log.Config{})

	debugLevel := &log.AllowedLevel{}
	if err := debugLevel.Set("debug"); err != nil {
		t.Fatal(err)
	}
	infoLevel := &log.AllowedLevel{}
	if err := infoLevel.Set("info"); err != nil {
		t.Fatal(err)
	}

	recorder := &recordKeyvalLogger{}
	logger.Base = recorder
	logger.SetLevel(debugLevel)
	if err := log.Debug(logger).Log("hello", "world"); err != nil {
		t.Fatal(err)
	}
	if recorder.count != 1 {
		t.Fatal("log not found")
	}

	recorder.count = 0
	logger.SetLevel(infoLevel)
	if err := log.Debug(logger).Log("hello", "world"); err != nil {
		t.Fatal(err)
	}
	if recorder.count != 0 {
		t.Fatal("log found")
	}
	if err := log.Info(logger).Log("hello", "world"); err != nil {
		t.Fatal(err)
	}
	if recorder.count != 1 {
		t.Fatal("log not found")
	}
	if err := log.Debug(logger).Log("hello", "world"); err != nil {
		t.Fatal(err)
	}
	if recorder.count != 1 {
		t.Fatal("extra log found")
	}
}
