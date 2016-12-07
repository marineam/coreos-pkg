// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package capnrus

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func TestLevelValueFlag(t *testing.T) {
	lvl := LevelValue(logrus.DebugLevel)
	set := flag.NewFlagSet("", flag.PanicOnError)
	set.Var(&lvl, "level", "")
	set.Parse([]string{"-level", "info"})
	if lvl.Level() != logrus.InfoLevel {
		t.Errorf("Unexpected level: %s", lvl)
	}
}

func TestLevelValuePflag(t *testing.T) {
	lvl := LevelValue(logrus.DebugLevel)
	set := pflag.NewFlagSet("", pflag.PanicOnError)
	set.Var(&lvl, "level", "")
	set.Parse([]string{"--level", "info"})
	if lvl.Level() != logrus.InfoLevel {
		t.Errorf("Unexpected level: %s", lvl)
	}
}

func TestParseLevel(t *testing.T) {
	levels := map[string]logrus.Level{
		"panic":    logrus.PanicLevel,
		"CRITICAL": logrus.FatalLevel,
		"C":        logrus.FatalLevel,
		"fatal":    logrus.FatalLevel,
		"ERROR":    logrus.ErrorLevel,
		"0":        logrus.ErrorLevel,
		"E":        logrus.ErrorLevel,
		"error":    logrus.ErrorLevel,
		"WARNING":  logrus.WarnLevel,
		"1":        logrus.WarnLevel,
		"W":        logrus.WarnLevel,
		"warn":     logrus.WarnLevel,
		"NOTICE":   logrus.InfoLevel,
		"2":        logrus.InfoLevel,
		"N":        logrus.InfoLevel,
		"INFO":     logrus.InfoLevel,
		"3":        logrus.InfoLevel,
		"I":        logrus.InfoLevel,
		"info":     logrus.InfoLevel,
		"DEBUG":    logrus.DebugLevel,
		"4":        logrus.DebugLevel,
		"D":        logrus.DebugLevel,
		"debug":    logrus.DebugLevel,
		"TRACE":    logrus.DebugLevel,
		"5":        logrus.DebugLevel,
		"T":        logrus.DebugLevel,
	}

	for s, v := range levels {
		if l, err := ParseLevel(s); err != nil {
			t.Errorf("Parsing %q failed: %v", s, err)
		} else if l != v {
			t.Errorf("Parsing %q returned %s, expected %s", s, l, v)
		}
	}
}
