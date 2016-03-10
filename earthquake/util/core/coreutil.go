// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package core provides version information, debug information, and the initializer
package core

import (
	"os"

	log "github.com/cihub/seelog"
	"github.com/osrg/earthquake/earthquake/explorepolicy"
	"github.com/osrg/earthquake/earthquake/signal"
	logutil "github.com/osrg/earthquake/earthquake/util/log"
)

const EarthquakeVersion = "0.2.0-SNAPSHOT"

// Returns true if EQ_DEBUG is set
func DebugMode() bool {
	return os.Getenv("EQ_DEBUG") != ""
}

// Initializes the Earthquake system
func Init() {
	debug := DebugMode()
	logutil.InitLog("", debug)
	signal.RegisterKnownSignals()
	explorepolicy.RegisterKnownExplorePolicies()
}

// Prints information on panic
// Also prints the stack trace if it is in the debug mode (EQ_DEBUG)
func Recoverer() {
	debug := DebugMode()
	if r := recover(); r != nil {
		log.Criticalf("PANIC: %s", r)
		if debug {
			panic(r)
		} else {
			log.Info("Hint: For debug info, please set \"EQ_DEBUG\" to 1.")
			os.Exit(1)
		}
	}
}