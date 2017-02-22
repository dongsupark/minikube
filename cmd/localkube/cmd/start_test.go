/*
Copyright 2017 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"testing"
	"time"
)

func TestStartAPIServer(t *testing.T) {
	lks := NewLocalkubeServer()

	lks.NewAPIServer().Start()

	// It is necessary to sleep, to make sure that SimpleServer.serverRoutine()
	// finishes. Without the sleep, Start() does not run serverRoutine() at all
	// and exits silently.
	time.Sleep(1 * time.Second)

	lks.NewAPIServer().Stop()
}
