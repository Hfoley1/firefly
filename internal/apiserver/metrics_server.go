// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

package apiserver

import (
	"github.com/hyperledger/firefly/internal/config"
)

const (
	MetricsEnabled = "enabled"
	MetricsPath    = "path"
)

func initMetricsConfPrefix(prefix config.Prefix) {
	prefix.AddKnownKey(MetricsEnabled, true)
	prefix.AddKnownKey(MetricsPath, "/metrics")
}