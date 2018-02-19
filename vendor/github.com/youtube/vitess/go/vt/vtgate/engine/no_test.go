/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package engine

import (
	"testing"
)

// This package is mostly tested from VTGate. This place-holder test
// is to document this fact. In order to make sure all code paths are
// covered, you need to run the coverage tool from the vtgate directory
// using the following command:
//	go test -coverprofile=c.out -coverpkg ./engine,.
// You can then follow it with:
//	go tool cover -html=c.out -o=c.html
// to see the html output.
func TestNothing(t *testing.T) {
}