// Copyright 2025 The Sigstore Authors
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

package server

import "time"

type HTTPConfig struct {
	host        string
	idleTimeout time.Duration
	port        int
}
type HTTPOption func(config *HTTPConfig)

func NewHTTPConfig(options ...func(config *HTTPConfig)) *HTTPConfig {
	config := &HTTPConfig{
		host:        "localhost",
		idleTimeout: 60 * time.Second,
		port:        8080,
	}
	for _, opt := range options {
		opt(config)
	}

	return config
}

func WithHTTPPort(port int) HTTPOption {
	return func(config *HTTPConfig) {
		config.port = port
	}
}

func WithHTTPHost(host string) HTTPOption {
	return func(config *HTTPConfig) {
		config.host = host
	}
}

func WithHTTPIdleTimeout(idleTimeout time.Duration) HTTPOption {
	return func(config *HTTPConfig) {
		config.idleTimeout = idleTimeout
	}
}
