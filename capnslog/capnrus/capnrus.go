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

// capnrus implements support for forwarding capnslog through a logrus Logger.
package capnrus

import (
	"log"

	"github.com/coreos/pkg/capnslog"
	"github.com/sirupsen/logrus"
)

// SetLevel sets the log level in capnslog and the logrus standard logger.
func SetLevel(lvl logrus.Level) {
	logrus.SetLevel(lvl)
	switch lvl {
	case logrus.PanicLevel:
		capnslog.SetGlobalLogLevel(capnslog.CRITICAL)
	case logrus.FatalLevel:
		capnslog.SetGlobalLogLevel(capnslog.CRITICAL)
	case logrus.ErrorLevel:
		capnslog.SetGlobalLogLevel(capnslog.ERROR)
	case logrus.WarnLevel:
		capnslog.SetGlobalLogLevel(capnslog.WARNING)
	case logrus.InfoLevel:
		capnslog.SetGlobalLogLevel(capnslog.INFO)
	case logrus.DebugLevel:
		capnslog.SetGlobalLogLevel(capnslog.TRACE) // Or DEBUG?
	default:
		panic("Unhandled loglevel")
	}
}

// UseStandardLogger redirects log and capnslog through the logrus standard logger.
func UseStandardLogger() {
	UseLogger(logrus.StandardLogger())
}

// UseLogger redirects log and capnslog through the given logrus Logger.
func UseLogger(logger *logrus.Logger) {
	capnslog.SetFormatter(&logrusAdapter{logger: logger})

	// Since capnslog will hijack log by default lets hijack again
	// to logrus to skip the extra trip through capnslog.
	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(logger.Writer())
}

// logrusAdapter is a capnslog.Formatter that logs through logrus
type logrusAdapter struct {
	logger *logrus.Logger
}

func (l *logrusAdapter) Format(pkg string, lvl capnslog.LogLevel, _ int, entries ...interface{}) {
	entry := l.logger.WithField("package", pkg)
	switch lvl {
	case capnslog.CRITICAL:
		// CRITICAL could be either Fatal or Panic, capnslog.PackageLogger
		// is the one to call Exit() or panic() after logging so use Error.
		entry.Error(entries...)
	case capnslog.ERROR:
		entry.Error(entries...)
	case capnslog.WARNING:
		entry.Warn(entries...)
	case capnslog.NOTICE:
		// NOTICE does not exist in logrus, squash to Info.
		entry.Info(entries...)
	case capnslog.INFO:
		entry.Info(entries...)
	case capnslog.DEBUG:
		entry.Debug(entries...)
	case capnslog.TRACE:
		// TRACE does not exist in logrus, squash to Debug.
		entry.Debug(entries...)
	default:
		panic("Unhandled loglevel")
	}
}

func (l *logrusAdapter) Flush() {}
