//
// Copyright 2025 The Sigstore Authors.
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

package client

import (
	"net/http"
)

type roundTripper struct {
	http.RoundTripper
	userAgent string
}

func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.userAgent != "" {
		req.Header.Set("User-Agent", rt.userAgent)
	}
	return rt.RoundTripper.RoundTrip(req)
}

func CreateRoundTripper(inner http.RoundTripper, userAgent string) http.RoundTripper {
	if inner == nil {
		inner = http.DefaultTransport
	}
	if userAgent == "" {
		return inner
	}
	return &roundTripper{
		RoundTripper: inner,
		userAgent:    userAgent,
	}
}
