// Copyright 2018 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package gogo_compat

import (
	"testing"

	"github.com/lo5/twirp/external/descriptors"
)

func TestCompilation(t *testing.T) {
	// Test passes if this package compiles
}

func TestReflection(t *testing.T) {
	// Despite use of gogo, we should still be able to reflect on the service.
	var s Svc
	server := NewSvcServer(s, nil)
	fd, sd, err := descriptors.ServiceDescriptor(server)
	if err != nil {
		t.Fatalf("ServiceDescriptor err: %v", err)
	}
	if have, want := fd.GetPackage(), "twirp.internal.twirptest.gogo_compat"; have != want {
		t.Errorf("bad package, have=%q, want=%q", have, want)
	}
	if have, want := sd.GetName(), "Svc"; have != want {
		t.Errorf("bad service name, have=%q, want=%q", have, want)
	}
}
