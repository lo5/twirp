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

package main

import (
	"context"

	"github.com/lo5/twirp/clientcompat/internal/clientcompat"
)

type clientCompat struct {
	method func(context.Context, *clientcompat.Req) (*clientcompat.Resp, error)
	noop   func(context.Context, *clientcompat.Empty) (*clientcompat.Empty, error)
}

func (c *clientCompat) Method(ctx context.Context, req *clientcompat.Req) (*clientcompat.Resp, error) {
	return c.method(ctx, req)
}

func (c *clientCompat) NoopMethod(ctx context.Context, e *clientcompat.Empty) (*clientcompat.Empty, error) {
	return c.noop(ctx, e)
}
