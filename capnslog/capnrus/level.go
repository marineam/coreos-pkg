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
	"github.com/coreos/pkg/capnslog"
	"github.com/sirupsen/logrus"
)

// LevelValue implements flag.Value and pflag.Value for logrus.Level
type LevelValue logrus.Level

func (l LevelValue) String() string {
	return l.Level().String()
}

func (l *LevelValue) Set(s string) error {
	lvl, err := ParseLevel(s)
	*l = LevelValue(lvl)
	return err
}

func (l *LevelValue) Type() string {
	return "string"
}

// l.Level() is shorthand for logrus.Level(l)
func (l LevelValue) Level() logrus.Level {
	return logrus.Level(l)
}

// ParseLevel translates both capnslog and logrus level strings their corresponding logrus level values.
func ParseLevel(s string) (logrus.Level, error) {
	if lvl, err := capnslog.ParseLevel(s); err == nil {
		switch lvl {
		case capnslog.CRITICAL:
			return logrus.FatalLevel, nil
		case capnslog.ERROR:
			return logrus.ErrorLevel, nil
		case capnslog.WARNING:
			return logrus.WarnLevel, nil
		case capnslog.NOTICE:
			return logrus.InfoLevel, nil
		case capnslog.INFO:
			return logrus.InfoLevel, nil
		case capnslog.DEBUG:
			return logrus.DebugLevel, nil
		case capnslog.TRACE:
			return logrus.DebugLevel, nil
		}
	}
	return logrus.ParseLevel(s)
}
